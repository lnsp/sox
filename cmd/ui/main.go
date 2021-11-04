package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
		jsonMachines := make([]jsonMachine, len(resp.Machines))
		for i := range resp.Machines {
			jsonMachines[i] = jsonMachine{
				ID:     resp.Machines[i].Id,
				Name:   resp.Machines[i].Name,
				Status: resp.Machines[i].Status.String(),
			}
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(jsonMachines)
	})
}

func (handler *APIHandler) listSSHKeys() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (handler *APIHandler) listImages() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
