package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
)

const (
	add = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(add, grpc.WithInsecure())

	if err != nil {
		log.Fatalln("disconnected -- ", err)
	}

	defer conn.Close()
	greetclient := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	res, err := greetclient.SayHello(ctx, &pb.HelloRequest{Name: "golang server"})

	if err != nil {
		log.Println("grpc error")
		os.Exit(1)
	}

	log.Println("response is:  ", res.Message)
}
