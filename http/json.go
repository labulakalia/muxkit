package http

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var protoM = protojson.MarshalOptions{
	UseProtoNames:   true,
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
}

func Marshal(m proto.Message) ([]byte, error) {
	return protoM.Marshal(m)
}

func UnMarshal(data []byte, m proto.Message) error {
	return protojson.Unmarshal(data, m)
}
