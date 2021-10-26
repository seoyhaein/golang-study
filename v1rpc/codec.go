package v1rpc

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
)

// https://sourcegraph.com/github.com/micro/micro/-/blob/service/client/grpc/codec.go?L31
type (
	protoCodec struct{}
	wrapCodec  struct{ encoding.Codec }
)

func (protoCodec) Marshal(v interface{}) ([]byte, error) {
	b, err := proto.Marshal(v.(proto.Message))
	// prometheus
	//sentBytes.Add(float64(len(b)))
	return b, err
}

func (protoCodec) Unmarshal(data []byte, v interface{}) error {
	//receivedBytes.Add(float64(len(data)))
	return proto.Unmarshal(data, v.(proto.Message))
}

func (protoCodec) Name() string {
	return "proto"
}
