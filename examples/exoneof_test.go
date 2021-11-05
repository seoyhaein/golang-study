package examples_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	pb "github.com/seoyhaein/golang-study/examples"
)

// 참고 링크
// https://software-factotum.medium.com/protobuf-and-go-handling-oneof-field-type-172ca780ec47

var str = "./test.txt"

/*
	마샬링, 언마살링을 쓸때는 import 를 잘해야 한다.

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"

	https://stackoverflow.com/questions/65563409/importing-proto-files-from-different-package-causes-missing-method-protoreflect
*/
func TestExoneofA(t *testing.T) {

	f, err := os.OpenFile(str, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("file open failed: \n %+v", err)
	}
	defer f.Close()

	// 여기서 둘중하나를 넣을 수도 있음. 안넣으면 실패로 뜨게 했음.
	// 메세지의 부모 자식의 관계는 '부모메세지_자식메세지' 로 표현됨.
	// 이 부분은 좀더 exscala.proto 를 통해서 자세히 다룸.
	op1 := &pb.Patch_Copy{Start: 10, End: 100}
	op2 := &pb.Patch_Insert{RawBytes: []byte{'a', 'b', 'c', 'd', 'e', 'a', 'b', 'c', 'd', 'e'}}

	// 질문 : 여기서 왜 Patch_CopyOp 과 Patch_InsertOp 이 들어 갈 수 있을까?
	// oneof 를 살펴보면 된다
	ops := []*pb.Patch{
		{Op: &pb.Patch_CopyOp{CopyOp: op1}},
		{Op: &pb.Patch_InsertOp{InsertOp: op2}},
	}
	// repeated 를 써서 슬라이스로 표현했다. []*Patch 타입
	patchins := &pb.Instructions{Operations: ops}

	// 이부분이 exany_test.go 와 달라지기 작하는 점이다.
	out, err := proto.Marshal(patchins)

	if err != nil {
		t.Fatalf("마샬 실패: \n %+v", err)
	}
	var n int = 0
	n, err = f.Write(out)

	if err != nil {
		t.Fatalf("파일 쓰기 실패: \n %+v", err)
	}
	t.Logf("사이즈 : \n %+v", n)
}

// TestExoneofA 를 먼저 실행시켜야 한다.
func TestExoneofB(t *testing.T) {
	in, err := ioutil.ReadFile(str)

	if err != nil {
		t.Fatalf("파일 읽기 실패: \n %+v", err)
	}
	// 이부분이 exany_test.go 와 달라지기 작하는 점이다.
	// 언마샬할때 받을 녀석
	instruction := &pb.Instructions{}

	if err := proto.Unmarshal(in, instruction); err != nil {
		t.Fatalf("언마샬링 실패: \n %+v", err)
	}

	for _, patch := range instruction.Operations {

		switch op := patch.Op.(type) {
		case *pb.Patch_CopyOp:
			t.Logf("Copy Operation start: %d, end : %d\n", op.CopyOp.Start, op.CopyOp.End)
		case *pb.Patch_InsertOp:
			t.Logf("Insert Operation Rawbytes length: %d\n", len(op.InsertOp.RawBytes))
		default:
			t.Fatal("매치 안됨.")
		}
	}
}
