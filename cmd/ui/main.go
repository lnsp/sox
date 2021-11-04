package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
		server.Handler = handlers.CORS()(server.Handler)
	}
	// Setup logging
	server.Handler = handlers.LoggingHandler(os.Stdout, server.Handler)
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
	mux.Handle("/", handler.version()).Methods("GET")
	mux.Handle("/machines", handler.listMachines()).Methods("GET")
	mux.Handle("/ssh-keys", handler.listSSHKeys()).Methods("GET")
	mux.Handle("/images", handler.listImages()).Methods("GET")
	mux.Handle("/networks", handler.listNetworks()).Methods("GET")
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
