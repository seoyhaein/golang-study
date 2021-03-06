package main

import (
	"context"
	"github.com/seoyhaein/golang-study/examples"
	"log"
	"os"
	"time"

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
	greetclient := examples.NewGreeterClient(conn)

	// server 서는 5초로 설정해 두었다. 5초 보다 느리게 time duration 을 두면 성공
	// 5초 보다 빠르면 실패한다.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := greetclient.SayHello(ctx, &examples.HelloRequest{Name: "golang server"})

	// 개선해보자 error 을 출력하는 방식으로
	if err != nil {
		log.Println("grpc error")
		os.Exit(1)
	}

	log.Println("response is:  ", res.Message)
}
