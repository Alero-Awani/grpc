// Package models provides the data models used in the application
package models

import (
	pb "github.com/aleroawani/grpc/pb/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Laptop represents a laptop with its associated attributes
type CreateLaptop struct {
Id 				string 								`bson:"uuid"`
	Brand			string							`bson:"brand"`
	Name			string							`bson:"name"`
	CPU				*pb.CPU							`bson:"cpu"`
	Ram				*pb.Memory						`bson:"ram"`
	GPUs			[]*pb.GPU						`bson:"gpus"`
	Storages		[]*pb.Storage					`bson:"storages"`
	Screen			*pb.Screen						`bson:"screen"`
	Keyboard		*pb.Keyboard					`bson:"keyboard"`
	Weight			float64							`bson:"weight"`
	Price			float64							`bson:"price"`
	Release_year	uint32							`bson:"release_year"`
	Updated_at		*timestamppb.Timestamp			`bson:"updated_at"` 
}

type NewCreateLaptop struct {
	Id        		primitive.ObjectID 				`bson:"_id,omitempty"`
	UUID 			string 							`bson:"uuid"`
	Brand			string							`bson:"brand"`
	Name			string							`bson:"name"`
	CPU				*pb.CPU							`bson:"cpu"`
	Ram				*pb.Memory						`bson:"ram"`
	GPUs			[]*pb.GPU						`bson:"gpus"`
	Storages		[]*pb.Storage					`bson:"storages"`
	Screen			*pb.Screen						`bson:"screen"`
	Keyboard		*pb.Keyboard					`bson:"keyboard"`
	Weight			float64							`bson:"weight"`
	Price			float64							`bson:"price"`
	Release_year	uint32							`bson:"release_year"`
	Updated_at		*timestamppb.Timestamp			`bson:"updated_at"` 
}


