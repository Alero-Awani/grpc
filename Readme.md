# Introduction

Target: Build a pc-book app to manage and search for laptop configurations.

## What is grpc 

`Remote Procedure Calls` - This a protocol that allows a program to execute a procedure of another program located in another computer without the developer explicitly coding the details for the remote interaction as it is automatically handled by the underlying framework.

On the backend side, we might have many services written in different languages like Go, python etc, following a microservice architecture.

In order to communicate with each other, they must all agree on a set of `API contracts` to exchange information including; 

- Communication channel: REST, SOAP, message queue
- Authentication mechanism: Basic, OAuth, JWT
- Payload format: JSON, XML, binary
- Data model 
- Error handling




Grpc offers code generation using `protocol Buffers`
The Protocol buffers are contained in the `proto file`. The protofile is basically the `API contract` which includes the services and the payload messages.

(The proto file is where you can define the schema for the data that is being sent.)

You can define things like 

- Expected fields 
- Required fields 
- Optional fields 
- Object types for these fields 

You also define the procedures that you expect to expose as the name RPC implies in the protofile you are defining what procedures you want to be callable remotely by other microservices.

If you want to have the code auto generated, you run the proto file against a compiler(protocol buffer compiler or protoc) and the output will be source code in the respective language(depending on the programming language, we have to tell the compiler to use the correct gRPC plugin for it), in this case Golang. This autogenerated code is called `stubs`.

The output(generated code) is an interface that creates the code for you that implements the object types that you outlined in your proto file.

## Installation Steps 

### Define the message 

To send messages, we have to define the message first. The protocol buffer file `processor_message.go` defines the messages and the services.

In this file, we define the `Request`, `Response` and the `Service`

In order to generate the server's code and the messages' encoding we use the `protoc-gen-go` which is the code generator.

```sh
protoc --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto 
```


`pb` is the path of the directory where you want to put your generated code.

`proto/*.proto` is the path of the proto file.

our proto file is located in the `proto` folder, so we tell protoc to look for it in that folder. With the `go_out` parameter. We tell protoc to use grpc plugins to generate GO code and store them in the pb folder.


### For protoc to run properly after install 

`vi ~/.zshrc`

```sh
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```

`source ~/.bash_profile`


### STEPS 

- Write protobuf message to binary file 
- Read protobuf message from binary file
- Write protobuf message to JSON file 
- Compare the size of the binary file and JSON file to see which one is smaller.


**Step 1** - Create a sample package to generate random data.

`randomBool` - function to generate random boolean for the backlit field 

The `serializer` package has functions to serialize the laptop objects to files.

**Step 2**
`WriteProtobufToBinaryFile()` - Used to write a protobuf message to a file in binary format. In our case, the message is the laptop object. 

We can use the proto.Message interface to make it more general.

`proto.Marshal` - serialise the message to binary. 
The binary is written in `laptop.bin` in the temp folder after the test runs.

**Step 3** 
After step 2, we write a function that converts that binary to a protobuf object and test it.
`ReadProtobufFromBinaryFile` Thus function reads back that binary file into a protobuf message object.

In the test
`laptop2 := &pb.Laptop{}` - We create a new laptop 2 object.

then we call the ReadPro... function to read the binary file data into that object.

`require.True(t.Proto.Equal(laptop1, laptop2))`
here we use `proto.Equal` method to check that laptop2 contains the same data as laptop1.

**Step 4** 
Now in order to actually see what was converted, we have to convert the protobuf message from step 3 to JSON

To check for test coverage, in the test file, click on `Run Package tests` then you can go back to the main file e.g `file.go` to the see which part of the code is covered.

### Comparing the JSON to binary file

On the terminal inside the tmp folder run `ls -l`
You will see the size of the json file is a lot bigger than that of the binary file. Hence we will save a lot of bandwith when using grpc instead of noraml JSON API.

