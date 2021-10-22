package v1rpc

import (
	"math"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

// 추후 아래 상수들에 대해서 설명하기

const (
	grpcOverheadBytes = 512 * 1024
	maxStreams        = math.MaxUint32
	maxSendBytes      = math.MaxInt32
)

func init() {
	// 추후 설명하자.
	// https://pkg.go.dev/google.golang.org/grpc#section-readme
	// https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md#using-a-codec
	// https://sourcegraph.com/github.com/asim/go-micro/-/blob/plugins/server/grpc/README.md
	encoding.RegisterCodec(wrapCodec{protoCodec{}})
}

// tls 는 일단 지금은 생략하자.
func Server(gopts ...grpc.ServerOption) *grpc.Server {

	return nil
}
