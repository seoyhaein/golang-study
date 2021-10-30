package v1rpc

import (
	"context"

	pb "github.com/seoyhaein/golang-study/protos"
	"google.golang.org/grpc"
)

// 10/30
// TODO api 와 package 를 달리 해서 헷갈리는 것을 방지 하자
// 추후 파일 이름 변경 taskmanager.go 너무 김.

type TaskManSrv struct {
	// 테스트 용
	pb.UnimplementedGreeterServer
}

// 추가 확장해나가자.
// 일단 테스트용으로 만들자.

func RegisterTM(service *grpc.Server) {
	pb.RegisterTaskManagerServer(service, newTaskManSrv())
}

func RegisterHello(service *grpc.Server) {
	pb.RegisterGreeterServer(service, newTaskManSrvHello())
}

func newTaskManSrv() pb.TaskManagerServer {
	return &TaskManSrv{}
}

func newTaskManSrvHello() pb.GreeterServer {
	return &TaskManSrv{}
}

// TODO 10/30  향후 추가하는 방향 고려
/*

	func NewTaskManSrv() pb.GreeterServer {
		return &TaskManSrv{
			//TaskManServer(),
		}
	}

	func TaskManServer() pb.TaskManagerServer {
		return nil
	}
*/

// api 정리 context는 사용하지 않는다. 고민좀 해보자.
// client 가 종료 되도 server 는 계속 구동 되는 방식으로 진행해야 한다. 따라서 context 는 사용하지 않는다.
// 서버에서 처리된 결과는 1차적으로 서버에서 저장되고 (여기선 파일로 저장)
// 2차적으로 처리 결과가 전송되고, 만일 client 가 접속 종료후 재접속 시에는 서버에 저장된 결과를 전송한다.
// 일단 여기서는 간단하게 저장은 파일로 저장하는 방식으로 취하자.

// grpc connection check 이 필요하다.
// 아직은 struct 를 이용하지 않는다.
func (*TaskManSrv) SendMessage(ctx context.Context, in *pb.InputMessage) (*pb.OutputMessage, error) {

	return &pb.OutputMessage{Output: "test string"}, nil
}

func (*TaskManSrv) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
