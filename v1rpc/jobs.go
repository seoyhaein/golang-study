package v1rpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
)

// 10/30
// TODO api 와 package 를 달리 해서 헷갈리는 것을 방지 하자

type JobManSrv struct {
	// api 가 작성되지 않으면 메세지 뿌린다.
	pb.UnimplementedLongLivedJobCallServer
	// req id 와 처리 결과를 저장한다.
	subscribers sync.Map
}

// map 에 저장하는 녀석, 향후 수정.
type sub struct {
	req_id int64
	// output 을 저장하는 녀석
	output any.Any
	status pb.JobsResponse_Status

	stream   pb.LongLivedJobCall_SubscribeServer // 이건 여기다 넣어야 할까? 고민중
	finished chan<- bool
}

// TODO 11/5 exeRunner 형태도 마음에 안든다. 어떻게 고칠까?
func RegisterJobsManSrv(service *grpc.Server) {
	pb.RegisterLongLivedJobCallServer(service, newJobsManSrv())
	//pb.RegisterLongLivedJobCallServer(service, newJobsManSrv1(exeRunner1))
}

// TODO 11/5
// 초기 설정 세팅 해주는 부분
// 아래 두 함수 비교해보자.
func newJobsManSrv() pb.LongLivedJobCallServer {
	j := new(JobManSrv)
	go j.exeRunner()
	return j
}

func newJobsManSrv1(f func(j *JobManSrv)) pb.LongLivedJobCallServer {
	j := new(JobManSrv)
	return func() pb.LongLivedJobCallServer {
		// TODO 고루틴 에러 처리 관련 추후 작성.
		go f(j)
		return j
	}()
}

// mesos state.go 에서 doSubscribe 함수 참고.
func (j *JobManSrv) Subscribe(in *pb.JobsRequest, s pb.LongLivedJobCall_SubscribeServer) error {

	log.Println("job request ID: ", in.JobReqId)

	fin := make(chan bool)

	// map 에 저장한다.
	j.subscribers.Store(in.JobReqId, sub{stream: s, finished: fin})

	ctx := s.Context()
	// Keep this scope alive because once this scope exits - the stream is closed
	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %d", in.JobReqId) // 일단!
			return nil
		case <-ctx.Done():
			log.Printf("Client ID %d has disconnected", in.JobReqId) // 일단!
			return nil
		}
	}
}

// Unsubscribe 를 client 에서 실행 전까지 exeRunner 를 계속 구동 시킨다.
func (j *JobManSrv) Unsubscribe(ctx context.Context, req *pb.JobsRequest) (*pb.JobsResponse, error) {
	v, ok := j.subscribers.Load(req.JobReqId)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %d", req.JobReqId)
	}
	sub, ok := v.(sub)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}
	select {
	case sub.finished <- true:
		log.Printf("Unsubscribed client: %d", req.JobReqId)
	default:
		// default 구문 주목.
	}
	j.subscribers.Delete(req.JobReqId)
	return &pb.JobsResponse{JobReqId: req.JobReqId}, nil
}

// TODO 11/4 일단 여기에는 간단한 id 를 넣어주는 것만.
// 이녀석은 별도의 스레드로 보내야 할까? 그래야 문제가 없을듯.
// Subscribe 에서 채널을 통해 데이터를 보내면, exeRunner 에서 받아서 처리하는 방식으로 하는 것은 어떨까?
// exec.Command 에서 처리 된 결과에 따라서...
// 일단 루프 구문으로 처리했음.

func (j *JobManSrv) exeRunner() {
	log.Println("Starting data generation")
	for {
		time.Sleep(time.Second)

		// A list of clients to unsubscribe in case of error
		var unsubscribe []int64

		// Range 함수는 재귀호출 구문을 사용한다. 정말 좋은 예이다.
		j.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(int64)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sub, ok := v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}
			// TODO 11/5 error prone
			// Send data over the gRPC stream to the client
			if err := sub.stream.Send(&pb.JobsResponse{JobReqId: id}); err != nil {
				log.Printf("Failed to send data to client: %v", err)
				select {
				case sub.finished <- true:
					log.Printf("Unsubscribed client: %d", id)
				default:
					// Default case is to avoid blocking in case client has already unsubscribed
				}
				// In case of error the client would re-subscribe so close the subscriber stream
				unsubscribe = append(unsubscribe, id)
			}
			return true
		})

		// Unsubscribe erroneous client streams
		for _, id := range unsubscribe {
			j.subscribers.Delete(id)
		}
	}
}

func exeRunner1(j *JobManSrv) {
	log.Println("Starting data generation")
	for {
		time.Sleep(time.Second)

		// A list of clients to unsubscribe in case of error
		var unsubscribe []int64

		// Range 함수는 재귀호출 구문을 사용한다. 정말 좋은 예이다.
		j.subscribers.Range(func(k, v interface{}) bool {
			id, ok := k.(int64)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sub, ok := v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}
			// Send data over the gRPC stream to the client
			if err := sub.stream.Send(&pb.JobsResponse{JobReqId: id}); err != nil {
				log.Printf("Failed to send data to client: %v", err)
				select {
				case sub.finished <- true:
					log.Printf("Unsubscribed client: %d", id)
				default:
					// Default case is to avoid blocking in case client has already unsubscribed
				}
				// In case of error the client would re-subscribe so close the subscriber stream
				unsubscribe = append(unsubscribe, id)
			}
			return true
		})

		// Unsubscribe erroneous client streams
		for _, id := range unsubscribe {
			j.subscribers.Delete(id)
		}
	}
}
