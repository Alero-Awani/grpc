package service

import (
	"context"
	"fmt"

	"github.com/aleroawani/grpc/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// // LaptopStore is an interface to store laptop
// type LaptopStore interface {
// 	// Save saves the laptop to the store
// 	Save(laptop *pb.Laptop) error
// 	// Find finds a laptop by ID
// 	Find(id string) (*pb.Laptop, error)
// }


// // InMemoryLaptopStore stores laptop in memory
// type InMemoryLaptopStore struct {
// 	mutex sync.RWMutex
// 	data map[string]*pb.Laptop
// }


// // NewInMemoryLaptopStore returns a new InmemoryLaptopStore 
// func NewInMemoryLaptopStore() *InMemoryLaptopStore {
// 	return &InMemoryLaptopStore{
// 		data: make(map[string]*pb.Laptop),
// 	}
// }

// // Save saves the laptop to the store 
// func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
// 	store.mutex.Lock()
// 	defer store.mutex.Unlock()

// 	if store.data[laptop.Id] != nil {
// 		return ErrAlreadyExists 
// 	}

// 	// deep copy 
// 	other := &pb.Laptop{}
// 	err := copier.Copy(other, laptop)
// 	if err != nil {
// 		return fmt.Errorf("cannot copy laptop data: %w", err)
// 	}

// 	store.data[other.Id] = other
// 	return nil
// }

// // Find finds a laptop by ID
// func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
// 	store.mutex.RLock()
// 	defer store.mutex.RUnlock()

// 	laptop := store.data[id]
// 	if laptop == nil {
// 		return nil, nil 
// 	}

// 	// deep copy 
// 	other := &pb.Laptop{}
// 	err := copier.Copy(other, laptop)
// 	if err != nil {
// 		return nil , fmt.Errorf("cannot copy laptop data: %w", err)
// 	}

// 	return other, nil

// }

// MongoDB

//Define an interface we must implement to perfrom the createlaptop operation
type LaptopService interface {
	CreateLaptop(*models.CreateLaptop) (*models.NewCreateLaptop, error)
}

// struct used to post data to the mongodb collection
type LaptopServiceImpl struct {
	postCollection	*mongo.Collection
	ctx				context.Context

}

func NewLaptopService(postCollection *mongo.Collection, ctx context.Context) LaptopService {
	return &LaptopServiceImpl{postCollection, ctx}
}

// create method to implement the LaptopService interface
func(l *LaptopServiceImpl) CreateLaptop(post *models.CreateLaptop) (*models.NewCreateLaptop, error) {
	//logic to save the data in mongodb
	res, err := l.postCollection.InsertOne(l.ctx, post)

	// check for potential errors 
	if err != nil {
		// return internal grpc error 
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	var newLaptop *models.NewCreateLaptop

	// use the objectId to retreive the newly created post 
	query := bson.M{"_id": res.InsertedID}
	if err = l.postCollection.FindOne(l.ctx, query).Decode(&newLaptop); err != nil {
		return nil, err

	}

	return newLaptop, err
}