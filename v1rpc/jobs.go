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

// TODO 11/5
// 참고 os/exec
// https://medium.com/rungo/executing-shell-commands-script-files-and-executables-in-go-894814f1c0f7

// 10/30
// TODO api 와 package 를 달리 해서 헷갈리는 것을 방지 하자
// client 로 부터 전송된 JobsRequest 데이터의 임시 저장 또는 필요 데이터를 메서드에 넘기는 역활
// 따라서 JobsRequest 일부 동일한 타입의 필드를 가질 수 있으며 추가적인 필드를 가질 수 있다.
type JobManSrv struct {
	// api 가 작성되지 않으면 메세지 뿌린다.
	pb.UnimplementedLongLivedJobCallServer
	// req id 와 처리 결과를 저장한다.
	subscribers sync.Map
}

// map 에 저장하는 녀석, 향후 수정.
// 음.....

type sub struct {
	req_id int64
	// output 을 저장하는 녀석
	output any.Any
	status pb.JobsResponse_Status

	stream   pb.LongLivedJobCall_SubscribeServer // 이건 여기다 넣어야 할까? 고민중
	finished chan<- bool
}

func RegisterJobsManSrv(service *grpc.Server) {
	pb.RegisterLongLivedJobCallServer(service, newJobsManSrv())
	//pb.RegisterLongLivedJobCallServer(service, newJobsManSrv1(exeRunner1))
}

// https://hwan-shell.tistory.com/339 인라인 함수, stack 에서 힙으로
// 초기 설정 세팅 해주는 부분
/*func newJobsManSrv() pb.LongLivedJobCallServer {
	j := new(JobManSrv)
	go j.exeRunner()
	return j
}

func newJobsManSrv1(f func(j *JobManSrv)) pb.LongLivedJobCallServer {
	j := new(JobManSrv)
	return func() pb.LongLivedJobCallServer {
		go f(j)
		return j
	}()
}*/

func newJobsManSrv() pb.LongLivedJobCallServer {
	j := new(JobManSrv)
	return j
}

// mesos state.go 에서 doSubscribe 함수 참고.
// TODO 11/6 이름이 맘에 안듬 추후 수정
func (j *JobManSrv) Subscribe(in *pb.JobsRequest, s pb.LongLivedJobCall_SubscribeServer) error {

	log.Println("job request ID: ", in.JobReqId)

	fin := make(chan bool)

	// map 에 저장한다.
	j.subscribers.Store(in.JobReqId, sub{stream: s, finished: fin})
	ctx := s.Context()

	// TODO 11/6 추후 수정 asap. 테스트 코드를 만들어서 진행 후 적용
	// 현재 고루틴 방식으로 만들어지지 않았다.
	go j.exeRunner()

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

// 새벽에 작업한것이라 일단 삭제
// 대충짬 졸리다. 아이디어가 떠올라서 그냥 메모 형태로 작성중. 아닌것 같다. 처음 생각했던것이랑 다른 거 같은데. TT
/*func (j *JobManSrv) SubscribeT(in *pb.JobsRequest, s pb.LongLivedJobCall_SubscribeServer) error {
	// to do something.
	// block part.
	return subscribe(exeRunner1,j)
}

func subscribe(f func(j *JobManSrv), j *JobManSrv) error {
	go f(j)
	return nil
}*/

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

// TODO 11/4 일단 여기에는 간단한 id 를 넣어주는 것만. 제일 먼저 처리 하자.(1,2,3,4 수정)
// 1. Subscribe 에서 채널을 통해 데이터를 보내면, exeRunner 에서 받아서 처리하는 방식으로 하는 것은 어떨까?
// 2. exec.Command 에서 처리 된 결과에 따라서...
// 3. 일단 루프 구문으로 처리했음.
// 4. 해당 함수는 고루틴 용으로 제작된 함수는 아니다. 하지만 일단 리턴값은 없기 때문에 그냥 둔다.

// 참고 : https://mingrammer.com/gobyexample/spawning-processes/
// https://medium.com/rungo/executing-shell-commands-script-files-and-executables-in-go-894814f1c0f7
// 고루틴 처리시 종료 시점은 결국 shell script 의 실행완료 이다. 즉, 해당 별도의 프로세스가 실행완료할때 해당 고루틴은 종료!
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

// 일단 삭제

/*func exeRunner1(j *JobManSrv) {
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
}*/
