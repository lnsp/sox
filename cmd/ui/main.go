package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/valar/virtm/api"
	"github.com/valar/virtm/meta"
	"github.com/valar/virtm/ui"
	"google.golang.org/grpc"
)

func main() {
	serveAddress := flag.String("serve", "localhost:9877", "Listen on address")
	virtmEndpoint := flag.String("virtm-endpoint", "localhost:9876", "VirtM server endpoint")
	virtmInsecure := flag.Bool("virtm-insecure", true, "Use insecure gRPC client")
	dev := flag.Bool("dev", false, "Disables static file hosting and relaxes CORS filtering")
	flag.Parse()

	// Connect to virtm
	virtm, err := connect(*virtmEndpoint, *virtmInsecure)
	if err != nil {
		log.Fatalln("connect to virtm:", err)
	}

	router := mux.NewRouter()
	// Setup API endpoints
	apiHandler := &APIHandler{
		Client: virtm,
	}
	if err := apiHandler.Init(router.PathPrefix("/api/").Subrouter()); err != nil {
		log.Fatalln("init api:", err)
	}
	// Setup static endpoints
	if !*dev {
		fs := http.FileServer(http.FS(ui.FS))
		router.PathPrefix("/").Handler(fs)
	}
	// Setup server and listen
	server := &http.Server{
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		Handler:      router,
		Addr:         *serveAddress,
	}
	// Setup CORS
	if *dev {
		methods := handlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST"})
		headers := handlers.AllowedHeaders([]string{"Content-Type"})
		origins := handlers.AllowedOrigins([]string{"*"})
		server.Handler = handlers.CORS(methods, headers, origins)(router)
	}
	// Listen and serve
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("serve:", err)
	}
}

type APIHandler struct {
	Client api.VirtMClient
}

func (handler *APIHandler) Init(mux *mux.Router) error {
	// Setup routes
	mux.Handle("/", handler.version()).Methods(http.MethodGet)
	mux.Handle("/machines", handler.listMachines()).Methods(http.MethodGet)
	mux.Handle("/machines", handler.createMachine()).Methods(http.MethodPost)
	mux.Handle("/machines/{id}", handler.showMachineDetails()).Methods(http.MethodGet)
	mux.Handle("/machines/{id}", handler.deleteMachine()).Methods(http.MethodDelete)
	mux.Handle("/machines/{id}/trigger", handler.triggerMachine()).Methods(http.MethodPost).Queries("event", "{event}")
	mux.Handle("/ssh-keys", handler.listSSHKeys()).Methods(http.MethodGet)
	mux.Handle("/images", handler.listImages()).Methods(http.MethodGet)
	mux.Handle("/networks", handler.listNetworks()).Methods(http.MethodGet)
	mux.Handle("/activities", handler.listActivities()).Methods(http.MethodGet)
	return nil
}

func (handler *APIHandler) version() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg, _ := json.Marshal(struct {
			Version string `json:"version"`
		}{
			Version: meta.Version,
		})
		w.Write(msg)
	})
}

func (handler *APIHandler) deleteMachine() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, err := handler.Client.DeleteMachine(r.Context(), &api.DeleteMachineRequest{
			Id: vars["id"],
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("delete machine:", err)
			return
		}
		fmt.Fprintln(w, "ok")
	})
}

func (handler *APIHandler) showMachineDetails() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		resp, err := handler.Client.GetMachineDetails(r.Context(), &api.GetMachineDetailsRequest{
			Id: vars["id"],
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("machine details:", err)
			return
		}
		type jsonNetworkInterface struct {
			IPv4      string `json:"ipv4"`
			IPv6      string `json:"ipv6"`
			NetworkID string `json:"networkId"`
		}
		body := struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Status string `json:"status"`

			ImageID           string                 `json:"imageId"`
			SSHKeyIds         []string               `json:"sshKeyIds"`
			NetworkInterfaces []jsonNetworkInterface `json:"networkInterfaces"`
		}{
			ID:                resp.Machine.Id,
			Name:              resp.Machine.Name,
			Status:            resp.Machine.Status.String(),
			ImageID:           resp.Machine.ImageId,
			SSHKeyIds:         resp.Machine.SshKeyIds,
			NetworkInterfaces: make([]jsonNetworkInterface, len(resp.Machine.Networks)),
		}
		for i := range resp.Machine.Networks {
			body.NetworkInterfaces[i] = jsonNetworkInterface{
				IPv4:      resp.Machine.Networks[i].IpV4,
				IPv6:      resp.Machine.Networks[i].IpV6,
				NetworkID: resp.Machine.Networks[i].NetworkId,
			}
		}
		json.NewEncoder(w).Encode(&body)
	})
}

func (handler *APIHandler) listMachines() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler.Client.ListMachines(r.Context(), &api.ListMachinesRequest{})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("fetch machines:", err)
			return
		}
		type jsonMachine struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Status string `json:"status"`
		}
		machines := make([]jsonMachine, len(resp.Machines))
		for i := range resp.Machines {
			machines[i] = jsonMachine{
				ID:     resp.Machines[i].Id,
				Name:   resp.Machines[i].Name,
				Status: resp.Machines[i].Status.String(),
			}
		}
		json.NewEncoder(w).Encode(machines)
	})
}

func (handler *APIHandler) listSSHKeys() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler.Client.ListSSHKeys(r.Context(), &api.ListSSHKeysRequest{})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("fetch ssh keys:", err)
			return
		}
		type jsonSSHKey struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Fingerprint string `json:"fingerprint"`
		}
		sshKeys := make([]jsonSSHKey, len(resp.Keys))
		for i := range resp.Keys {
			sshKeys[i] = jsonSSHKey{
				ID:          resp.Keys[i].Id,
				Name:        resp.Keys[i].Name,
				Fingerprint: strings.Fields(resp.Keys[i].Pubkey)[1][:16],
			}
		}
		json.NewEncoder(w).Encode(sshKeys)
	})
}

func (handler *APIHandler) listImages() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler.Client.ListImages(r.Context(), &api.ListImagesRequest{})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("fetch images:", err)
			return
		}
		type jsonImage struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			System string `json:"system"`
		}
		images := make([]jsonImage, len(resp.Images))
		for i := range resp.Images {
			images[i] = jsonImage{
				ID:     resp.Images[i].Id,
				Name:   resp.Images[i].Name,
				System: resp.Images[i].System.String(),
			}
		}
		json.NewEncoder(w).Encode(images)
	})
}

func (handler *APIHandler) listNetworks() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler.Client.ListNetworks(r.Context(), &api.ListNetworksRequest{})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("fetch networks:", err)
			return
		}
		type jsonNetwork struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			IPv4 string `json:"ipv4"`
			IPv6 string `json:"ipv6"`
		}
		networks := make([]jsonNetwork, len(resp.Networks))
		for i := range resp.Networks {
			networks[i] = jsonNetwork{
				ID:   resp.Networks[i].Id,
				Name: resp.Networks[i].Name,
				IPv4: resp.Networks[i].IpV4.Subnet,
				IPv6: resp.Networks[i].IpV6.Subnet,
			}
		}
		json.NewEncoder(w).Encode(networks)
	})
}

func (handler *APIHandler) createMachine() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Name  string `json:"name"`
			Specs struct {
				Cpus   int64 `json:"cpus"`
				Memory int64 `json:"memory"`
				Disk   int64 `json:"disk"`
			} `json:"specs"`
			Image    string   `json:"imageId"`
			SSHKeys  []string `json:"sshKeyIds"`
			Networks []string `json:"networkIds"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "could not decode body", http.StatusBadRequest)
			log.Println("decode request body:", err)
			return
		}
		resp, err := handler.Client.CreateMachine(r.Context(), &api.CreateMachineRequest{
			Name: body.Name,
			Specs: &api.Machine_Specs{
				Cpus:   (body.Specs.Cpus),
				Memory: (body.Specs.Memory),
				Disk:   (body.Specs.Disk),
			},
			ImageId:    body.Image,
			SshKeyIds:  body.SSHKeys,
			NetworkIds: body.Networks,
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("create machine:", err)
			return
		}
		json.NewEncoder(w).Encode(struct {
			Id string `json:"id"`
		}{
			resp.Id,
		})
	})
}

func (handler *APIHandler) triggerMachine() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		evt := api.TriggerMachineRequest_Event(api.TriggerMachineRequest_Event_value[vars["event"]])
		resp, err := handler.Client.TriggerMachine(r.Context(), &api.TriggerMachineRequest{
			Id:    id,
			Event: evt,
		})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("trigger machine event:", err)
			return
		}
		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: resp.Status.String(),
		})
	})
}

func (handler *APIHandler) listActivities() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := handler.Client.ListActivities(r.Context(), &api.ListActivitiesRequest{})
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("list activities:", err)
			return
		}
		type jsonActivity struct {
			Timestamp time.Time `json:"timestamp"`
			Subject   string    `json:"subject"`
			Type      string    `json:"type"`
		}
		activities := make([]jsonActivity, len(resp.Activities))
		for i := range resp.Activities {
			activities[i] = jsonActivity{
				Timestamp: resp.Activities[i].Timestamp.AsTime(),
				Type:      resp.Activities[i].Type.String(),
				Subject:   resp.Activities[i].Subject,
			}
		}
		json.NewEncoder(w).Encode(struct {
			Activities []jsonActivity `json:"activities"`
		}{
			Activities: activities,
		})
	})
}

func connect(endpoint string, insecure bool) (api.VirtMClient, error) {
	var grpcOpts []grpc.DialOption
	if insecure {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}
	grpcClient, err := grpc.Dial(endpoint, grpcOpts...)
	if err != nil {
		return nil, fmt.Errorf("dial endpoint: %w", err)
	}
	return api.NewVirtMClient(grpcClient), nil
}
