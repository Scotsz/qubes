package protocol

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Protocol interface {
	Marshal(msg proto.Message) ([]byte, error)
	Unmarshal(data []byte, msg proto.Message) error
}

type json struct {
	mopts  protojson.MarshalOptions
	unopts protojson.UnmarshalOptions
}

func NewJson() *json {
	return &json{
		mopts: protojson.MarshalOptions{
			UseEnumNumbers:  false,
			EmitUnpopulated: true,
		},
		unopts: protojson.UnmarshalOptions{
			AllowPartial:   false,
			DiscardUnknown: false,
			Resolver:       nil,
		},
	}
}

func (j json) Marshal(msg proto.Message) ([]byte, error) {
	return j.mopts.Marshal(msg)
}

func (j json) Unmarshal(data []byte, msg proto.Message) error {
	return j.unopts.Unmarshal(data, msg)
}

type protobuf struct {
}

func NewProtobuf() *protobuf {
	return &protobuf{}
}
func (p protobuf) Marshal(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

func (p protobuf) Unmarshal(data []byte, msg proto.Message) error {
	return proto.Unmarshal(data, msg)
}
