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

// TODO 10/30 api 와 package 를 달리 해서 헷갈리는 것을 방지 하자
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
}

func newJobsManSrv() pb.LongLivedJobCallServer {
	j := new(JobManSrv)
	return j
}

// mesos state.go 에서 doSubscribe 함수 참고.
// TODO 11/6 이름이 맘에 안듬 추후 수정
func (j *JobManSrv) Subscribe(in *pb.JobsRequest, s pb.LongLivedJobCall_SubscribeServer) error {

	fin := make(chan bool)

	// map 에 저장한다.
	j.subscribers.Store(in.JobReqId, sub{stream: s, finished: fin})
	ctx := s.Context()

	// TODO 11/6 추후 수정 asap. 테스트 코드를 만들어서 진행 후 적용
	cmd, r := j.scriptRunner(ctx, in)

	// 별도의 스레드로 실행해야  shell script 가 완료된후 시작하지 않는다.
	// 여기서 사용된 error 는 리턴 되지 않는다.
	go func(cmd *exec.Cmd) {
		if cmd != nil {
			if err := cmd.Start(); err != nil {
				log.Printf("Error starting Cmd: %v", err)
				return
			}
			if err := cmd.Wait(); err != nil {
				log.Printf("Error waiting for Cmd: %v", err)
				return
			}
		}
	}(cmd)

	// 이녀석도 별도의 스레드로 돌린다.
	go j.reply(r)

	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %d", in.JobReqId)
			return nil
		case <-ctx.Done():
			log.Printf("Client ID %d has disconnected", in.JobReqId)
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
	// LookPath 때문에 echo 라고 써도 됨.
	cmd := exec.CommandContext(ctx, "echo", in.InputMessage)

	// StdoutPipe 쓰면 Run 및 기타 Run 을 포함한 method 를 쓰면 에러난다.
	r, _ := cmd.StdoutPipe()

	return cmd, r
}

// 참고 : https://www.youtube.com/watch?v=Naonb2XD_2Q

/*
	finished <- true 가 되는 경우
	1. send 가 실패한 경우
    2. shell script 가 실패 한 경우
    3. 성공한 경우
*/

func (j *JobManSrv) reply(i io.Reader) {

	var (
		unsubscribe []int64
		fin         int32 = 1
		sb          sub
		id          int64
		ok          bool
	)

	scan := bufio.NewScanner(i)

	for {
		// fin = 0  success end, 1 not end, 2 failed
		if fin == 0 || fin == 2 {
			sb.finished <- true
			break
		}

		j.subscribers.Range(func(k, v interface{}) bool {
			id, ok = k.(int64)
			if !ok {
				log.Printf("Failed to cast subscriber key: %T", k)
				return false
			}
			sb, ok = v.(sub)
			if !ok {
				log.Printf("Failed to cast subscriber value: %T", v)
				return false
			}

			b := scan.Scan()
			s := scan.Text()
			// TODO 11/20 error prone scan.Err() nil 안되는 이유 찾아보아야 함.
			if b != true {
				if scan.Err() == nil {
					fin = 0
					return false
				}
				log.Println(scan.Err())
				fin = 2
				return false
			}

			if err := sb.stream.Send(&pb.JobsResponse{JobResId: id, OutputMessage: s}); err != nil {
				log.Printf("Failed to send data to client: %v", err)

				select {
				case sb.finished <- true:
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
