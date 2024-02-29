package service

import (
	"context"
	"errors"
	"log"

	"github.com/aleroawani/grpc/models"
	pb "github.com/aleroawani/grpc/pb/proto"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	postCollection		*mongo.Collection
	LaptopService		LaptopService
}

// NewlaptopServer returns a new LaptopServer
func NewLaptopServer(postCollection *mongo.Collection, laptopService LaptopService) *LaptopServer {
	return &LaptopServer{
		postCollection,
		laptopService}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context, 
	req *pb.CreateLaptopRequest,) (*pb.CreateLaptopResponse, error) {

	laptop := req.GetLaptop()

	log.Printf("Received a create-laptop request with id: %s", laptop.Id)
	
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
	// save the laptop to database 
	data := &models.CreateLaptop{
		Id: 			laptop.GetId(),
		Brand:  		laptop.GetBrand(),		
		Name:			laptop.GetName(),			
		CPU: 			laptop.GetCpu(),			
		Ram:			laptop.GetRam(),			
		GPUs:			laptop.GetGpus(),		
		Storages:		laptop.GetStorages(),		
		Screen:			laptop.GetScreen(),	
		Keyboard: 		laptop.GetKeyboard(),	
		Weight:			laptop.GetWeightKg(),
		Price: 			laptop.GetPriceUsd(),		
		Release_year: 	laptop.GetReleaseYear(),	
		Updated_at: 	laptop.GetUpdatedAt(),	
	}

	newdata, err := server.LaptopService.CreateLaptop(data)

	code := codes.Internal
	if errors.Is(err, ErrAlreadyExists) {
		code = codes.AlreadyExists
	}
	if err != nil {
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}
	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: newdata.UUID,
	}
	return res, nil 
}


