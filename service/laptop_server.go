package service

import (
	"context"
	"log"

	pb "github.com/aleroawani/grpc/pb/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	Store LaptopStore

}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context, 
	req *pb.CreateLaptopRequest,) (*pb.CreateLaptopResponse, error) {

	laptop := req.GetLaptop()
	log.Printf("Received a create-laptop request with id: %s", laptop)
	
	if len(laptop.Id) > 0 {
		// Check if its a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	// save the laptop to store

}