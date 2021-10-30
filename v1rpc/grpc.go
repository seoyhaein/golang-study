package v1rpc

import (
	"log"
	"math"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

// 추후 아래 상수들에 대해서 설명하기
const (
	maxRequestBytes   = 1.5 * 1024 * 1024
	grpcOverheadBytes = 512 * 1024
	maxStreams        = math.MaxUint32
	maxSendBytes      = math.MaxInt32
)

// 일단 여기서만 쓰는 변수 선언 향후 구조체로 통합
var (
	lis     net.Listener
	network string
	address string
)

func init() {
	// 추후 설명하자.
	// https://pkg.go.dev/google.golang.org/grpc#section-readme
	// https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md#using-a-codec
	// https://sourcegraph.com/github.com/asim/go-micro/-/blob/plugins/server/grpc/README.md

	// 일단 테스트용으로 넣어둠.
	encoding.RegisterCodec(wrapCodec{protoCodec{}})
}

// tls 는 일단 지금은 생략하자.
// 추후 고쳐나가자. 일단은 가장 간단한 형태로 만들어 놓고 테스트 진행하면서 고도화 하자.
// gopts ...grpc.ServerOption 집어 넣는 방향
func Server() error {

	var err error = nil
	network = "tcp"
	address = ":50052"

	lis, err = net.Listen(network, address)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.MaxRecvMsgSize(int(maxRequestBytes+grpcOverheadBytes)))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))
	opts = append(opts, grpc.MaxConcurrentStreams(maxStreams))
	grpcServer := grpc.NewServer(opts...)

	// 일단 이걸 교체해야함.
	// 테스트 용으로 hello world 이용함.
	// 10/30
	//TODO api 와 package 를 달리 해서 헷갈리는 것을 방지 하자
	RegisterHello(grpcServer)

	log.Println("gRPC server started, address : ", address, "listener-address", lis.Addr().String())

	err = grpcServer.Serve(lis)
	if err != nil && strings.Contains(err.Error(), "use of closed network connection") {
		log.Println("gRPC server is shut down, address : ", address)

	} else {
		log.Println("gRPC server returned with error , address :", address)
	}
	return err
}
