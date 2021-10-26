package main

import (
	"flag"
	"log"
	"net"

	"github.com/valar/virtm/api"
	"github.com/valar/virtm/driver"
	"google.golang.org/grpc"
)

func main() {
	db := flag.String("db", "virtm.db", "path to database")
	address := flag.String("address", "localhost:9876", "address to listen on")
	flag.Parse()

	// start vm manager
	driver, err := driver.New(*db)
	if err != nil {
		log.Fatalf("failed to start driver: %v", err)
	}
	// setup listener
	listener, err := net.Listen("tcp", *address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// setup options
	var opts []grpc.ServerOption
	// start server
	grpcServer := grpc.NewServer(opts...)
	api.RegisterVirtMServer(grpcServer, driver)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
