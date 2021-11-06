package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
)

func main() {
	// Create multiple clients and start receiving data
	var (
		// wg     sync.WaitGroup
		client *longlivedClient
		err    error
	)

	/*	for i := 1; i <= 10; i++ {
		wg.Add(1)
		client, err = mkLonglivedClient(int64(i))
		if err != nil {
			log.Fatal(err)
		}
		// Dispatch client goroutine
		go client.start()
		time.Sleep(time.Second * 2)
	}*/

	client, err = mkLonglivedClient(int64(255))
	if err != nil {
		log.Fatal(err)
	}
	end := client.run()
	// 메세지가 더이상 없을때. 일단 주석 처리
	if end == io.EOF {
		/*err := client.unsubscribe()

		if err != nil {
			log.Fatalln(err)
		}*/

		// 정상 적인 종료
		os.Exit(0)
	}

	// 일단 비정상 종료로 처리하자.
	os.Exit(1)

	// The wait group purpose is to avoid exiting, the clients do not exit
	//wg.Wait()
	// TODO 11/5 error prone. check first asap!
	// 아마 아래 코드는 절대로 실행되지 않을 것임.
	// client.unsubscribe()
}

// longlivedClient holds the long lived gRPC client fields
type longlivedClient struct {
	client pb.LongLivedJobCallClient // client is the long lived gRPC client

	conn *grpc.ClientConn // conn is the client gRPC connection
	id   int64            // id is the client ID used for subscribing

	command *any.Any
}

// mkLonglivedClient creates a new client instance
// TODO 11/6 fix or make new a method.
// 리턴값을 JobsRequest 로 한다.
func mkLonglivedClient(id int64) (*longlivedClient, error) {
	conn, err := mkConnection()
	if err != nil {
		return nil, err
	}
	// TODO 11/6 일단 any.Any 테스트 하기 위해 그냥 넣어둠. 참고용으로 향후 메모 후 삭제함.
	// proto 파일도 수정해야함.

	comm := &pb.AnyString{
		Command: "echo 'hello world'",
	}
	b_comm, err := proto.Marshal(comm)
	if err != nil {
		return nil, err
	}

	return &longlivedClient{
		client: pb.NewLongLivedJobCallClient(conn),
		conn:   conn,
		id:     id,
		command: &any.Any{
			TypeUrl: "github.com/seoyhaein/golang-study",
			Value:   b_comm,
		},
	}, nil
}

// close is not used but is here as an example of how to close the gRPC client connection
func (c *longlivedClient) close() {
	if err := c.conn.Close(); err != nil {
		log.Fatal(err)
	}
}

// TODO 11/6 fix a method
// input param will be *pb.JobsRequest
// subscribe subscribes to messages from the gRPC server
// func (c *longlivedClient) subscribe(req *pb.JobsRequest) (pb.LongLivedJobCall_SubscribeClient, error) {
func (c *longlivedClient) subscribe() (pb.LongLivedJobCall_SubscribeClient, error) {
	log.Printf("Subscribing client ID: %d", c.id)
	return c.client.Subscribe(context.Background(), &pb.JobsRequest{JobReqId: c.id})
}

// unsubscribe unsubscribes to messages from the gRPC server
func (c *longlivedClient) unsubscribe() error {
	log.Printf("Unsubscribing client ID %d", c.id)
	_, err := c.client.Unsubscribe(context.Background(), &pb.JobsRequest{JobReqId: c.id})
	return err
}

// TODO 11/5 error prone
func (c *longlivedClient) start() {
	var err error
	// stream is the client side of the RPC stream
	var stream pb.LongLivedJobCall_SubscribeClient
	for {
		if stream == nil { // EOF ??
			if stream, err = c.subscribe(); err != nil {
				log.Printf("Failed to subscribe: %v", err)
				c.sleep()
				// Retry on failure
				continue
			}
		}
		// 데이터를 받고,
		response, err := stream.Recv()
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			// Clearing the stream will force the client to resubscribe on next iteration
			stream = nil
			c.sleep()
			// Retry on failure
			continue
		}

		log.Printf("Client ID %d got response: %d", c.id, response.JobReqId)
	}
}

func (c *longlivedClient) run() error {

	stream, err := c.subscribe()
	if err != nil {
		log.Printf("Failed to subscribe: %v", err)
		return err
	}

	for {
		response, err := stream.Recv()
		// 리턴할 메세지가 없으면..
		if err == io.EOF {
			return err
		}
		log.Println("Response : ", response.JobReqId)
		// 서버에서 데이터는 계속 보낸다. 이때 unscribe 를 할경우.
		c.unsubscribe()
	}
}

// sleep is used to give the server time to unsubscribe the client and reset the stream
func (c *longlivedClient) sleep() {
	time.Sleep(time.Second * 5)
}

func mkConnection() (*grpc.ClientConn, error) {
	return grpc.Dial("127.0.0.1:50052", []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
}