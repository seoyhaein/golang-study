package v1rpc

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"

	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
)

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

type sub struct {
	id       int64
	status   pb.JobsResponse_Status
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

	log.Println("1. job request ID: ", in.JobReqId)

	fin := make(chan bool)

	// map 에 저장한다.
	j.subscribers.Store(in.JobReqId, sub{stream: s, finished: fin})
	ctx := s.Context()

	// TODO 11/6 추후 수정 asap. 테스트 코드를 만들어서 진행 후 적용
	cmd, r := j.scriptRunner(ctx, in)
	log.Println("2. Run scriptRunner ")
	go j.reply(r)
	log.Println("3. Reply... ")

	cmd.Start()
	cmd.Wait()
	log.Println("4. Ready... ")
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
	return &pb.JobsResponse{JobResId: req.JobReqId}, nil
}

// stdio 에 바로 넣음. 아직 context 처리는 하지 않음.
func (j *JobManSrv) scriptRunner(ctx context.Context, in *pb.JobsRequest) (*exec.Cmd, io.Reader) {
	// script is InputMessage
	cmd := exec.CommandContext(ctx, in.InputMessage)

	// StdoutPipe 쓰면 Run 및 기타 Run 을 포함한 method 를 쓰면 에러난다.
	r, _ := cmd.StdoutPipe()

	return cmd, r
}

// TODO 11/13 버그 있음 해결 해야 함.
func (j *JobManSrv) reply(i io.Reader) {
	var unsubscribe []int64

	scan := bufio.NewScanner(i)
	// 디버깅 해야함.
	for scan.Scan() {
		s := scan.Text()
		// 주석처리 요망
		log.Println(s)
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

			if err := sub.stream.Send(&pb.JobsResponse{JobResId: id, OutputMessage: s}); err != nil {
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
