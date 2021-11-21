package v1rpc

// 패키지를 잘봐야 한다. 왜 패키지를 v1rpc 로 했을까?
// v1rpc_test 로 하지 않은 까닭은???

// 참고 : https://stackoverflow.com/questions/42102496/testing-a-grpc-service
//       http://www.inanzzz.com/index.php/post/w9qr/unit-testing-golang-grpc-client-and-server-application-with-bufconn-package

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var li *bufconn.Listener

func init() {
	li = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	// 이 메서드도 test 해야 하지만 스킵한다.
	RegisterJobsManSrv(s)

	// TODO 11/21 server 가 s.Stop() 되는지도 한번 확인해보자.
	go func() {
		if err := s.Serve(li); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return li.Dial()
}

func TestJobManSrv_Subscribe(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewLongLivedJobCallClient(conn)
	// TODO shell script 담는 부분 수정될때 변경해야함.
	resp, err := client.Subscribe(ctx, &pb.JobsRequest{JobReqId: 1, InputMessage: "hello world"})
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}
	// stream 이라서 루프 구문
	for {
		// 데이터를 받고,
		response, err := resp.Recv()

		if err == io.EOF {
			log.Println("EOF")
			break
		}
		if err != nil {
			// Clearing the stream will force the client to resubscribe on next iteration
			t.Fatalf("Failed to receive message : %v", err)
		}

		log.Println("value1 : ", response.JobResId, "value2 : ", response.OutputMessage)
	}
}
