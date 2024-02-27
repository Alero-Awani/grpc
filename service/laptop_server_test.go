package service_test

import (
	"testing"

	pb "github.com/aleroawani/grpc/pb/proto"
	"github.com/aleroawani/grpc/sample"
	"github.com/aleroawani/grpc/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func testServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID	:= sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)



	testCases := []struct {
		name	string
		laptop  *pb.Laptop
		store	service.LaptopStore
		code    codes.Code
	}{
		{
			name: 		"success_with_id",
			laptop:		sample.NewLaptop(),
			store: 		service.NewInMemoryLaptopStore(),
			code:		codes.OK,
		},
		{
			name:		"success-no_id",
			laptop:		laptopNoID,
			store: 		service.NewInMemoryLaptopStore(),
			code:		codes.OK,
		},

		{
			name:		"failure_invalid_id",
			laptop:		laptopInvalidID,
			store:		service.NewInMemoryLaptopStore(),
			code:		codes.InvalidArgument,			

		},
		{
			name:		"failure_duplicate_id",
			laptop:		laptopDuplicateID,
			store:		storeDuplicateID,
			code:		codes.AlreadyExists,
		},
	}
}


