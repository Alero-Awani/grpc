package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/aleroawani/grpc/pb/proto"
	"github.com/aleroawani/grpc/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

var (
	postCollection		*mongo.Collection
	laptopService 		service.LaptopService
	ctx					context.Context
	mongoclient			*mongo.Client
)

// func mongoConnect(){

// }

func main(){
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017/")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// collection
	postCollection = mongoclient.Database("laptop_proto_db").Collection("proto")
	laptopService = service.NewLaptopService(postCollection, ctx)


	// start server
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("starting server on port %d", *port)

	laptopServer := service.NewLaptopServer(postCollection, laptopService)
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

