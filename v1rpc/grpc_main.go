package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

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
	log.Println("Received: ", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
