package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

// 10/20
// server 에 대한 것도 향후 개선해야한다.
// protos/greet.pb.go 에 있다.
type server struct{ pb.UnimplementedGreeterServer }

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		// 개선점이 무엇일까?
		log.Println("net error")
		os.Exit(1)
	}

	// grpc part
	// grpc struct create
	s := grpc.NewServer()

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	// https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(s)
	pb.RegisterGreeterServer(s, &server{})

	// grpc server 시작
	log.Println("GRPC Listening Port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Println("Error")
		os.Exit(1)
		// log.Fatalln(err)
	}
}

// protos/greet.pb.go 에 원형이 존재하고 여기서 함수를 구현해주었다.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 5초 동안 sleep
	time.Sleep(time.Second * 5)

	select {
	case <-ctx.Done():
		log.Println("SayHello exit") // prints "context deadline exceeded"
		return &pb.HelloReply{Message: "SayHello exit"}, nil
	default:
	}

	log.Println("Received: ", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

/*
	여기서 개선점들이 무엇이 있을까?
	결국 API 들을 main 에 몰아 넣을 수 없고 따로 떼내어야 하고, 그리고 grpc part 에서 4개의 함수를 잘 구성하는 것이 중요하다.

	NewServer, RegisterGreeterServer (RegisterService), Server, Serveroptions
*/
