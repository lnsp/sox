package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/valar/virtm/api"
	"github.com/valar/virtm/driver"
	"github.com/valar/virtm/meta"
	"google.golang.org/grpc"
)

type config struct {
	Database struct {
		DSN string
	}
	Libvirt struct {
		URI string
	}
	Grpc struct {
		Address string
	}
}

var rootCmd = cobra.Command{
	Use:     "virtm [config]",
	Short:   "Experimental virtual machine manager",
	Version: meta.Version,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		run(args[0])
	},
}

func run(cfgpath string) {
	// load config
	cfgdata, err := os.ReadFile(cfgpath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// decode config
	var cfg config
	if err := toml.Unmarshal(cfgdata, &cfg); err != nil {
		log.Fatalf("failed to decode config: %v", err)
	}
	// start vm manager
	driver, err := driver.New(cfg.Database.DSN, cfg.Libvirt.URI)
	if err != nil {
		log.Fatalf("failed to start driver: %v", err)
	}
	log.Println("initialized vm driver")
	// setup listener
	listener, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening on", listener.Addr())
	// setup options
	var opts []grpc.ServerOption
	// start server
	grpcServer := grpc.NewServer(opts...)
	// handle shutdown signal
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-interrupts
		log.Println("received interrupt, stopping grpc server")
		grpcServer.GracefulStop()
	}()
	// register and start serving
	api.RegisterVirtMServer(grpcServer, driver)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("grpc server shutted down")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
