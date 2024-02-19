# Introduction

## Installation Steps 

### Define the message 

To send messages, we have to define the messaged first. The protocol buffer file `processor_message.go` defines the messages and the services.

In this file, we define the `Request`, `Response` and the `Service`

In order to generate the server's code and the messages' encoding we use the `protoc-gen-go` which is the code generator.

```sh
protoc --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto 
```


`pb` is the path of the directory where you want to put your generated code.

`proto/*.proto` is the path of the proto file.

our proto file is located in the `proto` folder, so we tell protoc to look for it in that folder. With the `go_out` parameter. We tell protoc to use grpc plugins to generate GO code and store them in the pb folder.


