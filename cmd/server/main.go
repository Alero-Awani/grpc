package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/aleroawani/grpc/pb/proto"
	"github.com/aleroawani/grpc/service"
	"google.golang.org/grpc"
)

func main(){
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("starting server on port %d", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	

}