### Steps 

1. **Define a proto service** that contains a unary gRPC to create a laptop

2. **Implement server in Go** Write server side code to handle the unary RPC request and save laptop to an in-memory storage

3. **Implement client in Go** Write client-side code to call the unary RPC on server. Also write unit test for the interaction between client and server.

4. **Handle errors and deadline** Learn how to set deadline for the request, check for cancelation, handle errors, and return suitable status code to the client.

## What is Unary grpc 

Unary RPC calls are the basic Request/ Response that we are familiar with 

![Unary API](assets/UnaryAPI.png)

- The client will send a message to the server and will receive one response from the server 
- Unary RPC calls are the most common for APIS 
- in gRPC, Unary Calls are defined using Protocol Buffers 

After creating the `Laptop Request`, `Laptop Response` and `Service` in the `laptop_service.proto` file. Run `make gen` to generate the stub code for the proto file.

In `laptop_service.pb.go`,

```go
func (x *CreateLaptopRequest) GetLaptop()
```
`createLaptopRequest` struct has a Function to get the input laptop

```go 
func (x *CreateLaptopResponse) GetId() string 

```
`createLaptopResponse` struct has a function to get the output ID of the laptop

In `laptop_service_grpc.pb.go` we have the `laptop service client interface`.
It is an interface because it will allow us implement our own custom client.E.g the mock client that will be used for unit testing.

```go 
type LaptopServiceClient interface
```

the `laptopServiceClient` struct implements the interface

Then we have the `laptopServiceServer` interface that has no implementation. This is because we are supposed to implement it ourselves. 
The interface has the `Createlaptop` function which has to satisfied by the laptopServer struct in `laptopServer.go`


The implementation is found in `laptop_server.go` file.


```go
func RegisterLaptopServiceServer(s grpc.ServiceRegistrar, srv LaptopServiceServer) {
	s.RegisterService(&LaptopService_ServiceDesc, srv)
}
```

The function above registers the Service on a specific grpc server, so that it can receive and handle requests from client.


## PROCEDURE 

- The `laptop_service.proto` file outlines a grpc service for creating laptops 
- `CreateLaptopRequest` is a message that *carries the details of a laptop to be created*.
- `CreateLaptopResponse` is a message that contains the identifier (id) of the newly created laptop.

- `service LaptopService` This is the service definition, and it contains the CreateLaptop RPC method. The method takes a CreateLaptopRequest as input and returns a CreateLaptopResponse.


### Server Implementation

- In `laptop_server.go` the server-side logic is implemented by creating a **Go struct that embeds the auto-generated server interface**(laptopServiceServer Interface). This struct(LaptopServer) is supposed to provide implementations for each of the service methods defined in the `laptop_service.proto` file. In this case we have the `Createlaptop` method

-  This function encapsulates the logic for creating and Initialising a laptopserver(dedicated function for creating instances)- it is a **constructor function**

```go
func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}
```

- `laptop := req.GetLaptop()` - This is to get the laptop object from the request.

- `if len(laptop.Id) > 0`If the client has already generated the laptop Id, check if it is a valid uuid or not.

- summary- we want to  create a new laptop ,and to do that we need to get the details for creating from the request, which is why we have req.GetLaptop(). Then we do somechecks on the given id or generate a new one. Then we are supposed to save to a db. But instead we will use an in memory store.

- `data map[string]*pb.Laptop` we use a map to store the data, where the key is the laptop id and the value is the laptop object.