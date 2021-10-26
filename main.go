package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/valar/virtm/api"
	"github.com/valar/virtm/driver"
	"google.golang.org/grpc"
)

func main() {
	db := flag.String("db", "virtm.db", "path to database")
	address := flag.String("address", "localhost:9876", "address to listen on")
	network := flag.String("libvirt", "qemu:///system", "libvirt instance to use")
	flag.Parse()

	// start vm manager
	driver, err := driver.New(*db, *network)
	if err != nil {
		log.Fatalf("failed to start driver: %v", err)
	}
	log.Println("initialized vm driver")
	// setup listener
	listener, err := net.Listen("tcp", *address)
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
