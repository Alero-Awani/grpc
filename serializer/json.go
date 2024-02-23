package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON convets protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers: false,
		EmitUnpopulated: true,
		Indent: "  ",
		UseProtoNames: true,
	}
	data, err := marshaler.Marshal(message)
	return string(data), err
}