package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/valar/virtm/api"
	"github.com/valar/virtm/driver"
	"github.com/valar/virtm/meta"
	"google.golang.org/grpc"
)

var rootCmd = cobra.Command{
	Use:     "virtm",
	Short:   "Experimental virtual machine manager",
	Version: meta.Version,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var db string
var address string
var libvirt string

func init() {
	flags := rootCmd.Flags()
	flags.StringVar(&db, "db", "virtm.db", "Path to database")
	flags.StringVar(&address, "address", "localhost:9876", "Address to listen on")
	flags.StringVar(&libvirt, "libvirt", "qemu:///system", "Libvirt instance")
}

func run() {
	// start vm manager
	driver, err := driver.New(db, libvirt)
	if err != nil {
		log.Fatalf("failed to start driver: %v", err)
	}
	log.Println("initialized vm driver")
	// setup listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
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
