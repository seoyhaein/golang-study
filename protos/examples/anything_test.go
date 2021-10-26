package examples_test

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"

	anything "github.com/seoyhaein/golang-study/protos/examples"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

/*
	import 된 패키지들을 잘보면 github 에서 추가된 녀석들과
	google.golang 에서 추가된 protobuf 패키지의 용도들이 다른 것을 확인 할 수 있다.
*/

func TestAnything(t *testing.T) {

	t1 := &timestamp.Timestamp{
		Seconds: 5, // easy to verify
		Nanos:   6, // easy to verify
	}

	serialized, err := proto.Marshal(t1)
	if err != nil {
		t.Fatal("could not serialize timestamp")
	}

	a := &anything.AnythingForYou{
		Anything: &any.Any{
			TypeUrl: "github.com/seoyhaein/",
			Value:   serialized,
		},
	}
	// 마샬링을 두가지 방식으로 한다. 성능 비교는 해봐야 알것 같다
	serializedA, err := proto.Marshal(a.Anything)
	serializedB, err1 := anypb.New(a.Anything)

	if err != nil {
		t.Fatal("proto.Marshal: could not serialize anything")
	}

	if err1 != nil {
		t.Fatal("anypb.New:  could not serialize anything")
	}
	/*
		언마샬 했을 때 담을 녀석들. a2, a3
	*/
	var a2 = &anything.AnythingForYou{
		Anything: &any.Any{},
	}

	var a3 = &anything.AnythingForYou{
		Anything: &any.Any{},
	}
	/*
		두가지 방식으로 언마샬 한다.
	*/
	if err := proto.Unmarshal(serializedA, a3.Anything); err != nil {
		t.Fatal("could not deserialize anything")
	}
	// UnmarshalTo,unmarahlFrom, New 등은 anypb 에 있다.
	// https://pkg.go.dev/google.golang.org/protobuf/types/known/anypb#UnmarshalTo
	/*
		UnmarshalTo unmarshals the underlying message from src into dst using the provided unmarshal options. It reports an error if dst is not of the right message type.

		If no options are specified, call src.UnmarshalTo instead.
	*/

	if err2 := serializedB.UnmarshalTo(a2.Anything); err2 != nil {
		t.Fatal("could not deserialize anything")
	}

	if a.Anything.TypeUrl != a3.Anything.TypeUrl {
		if !bytes.Equal(a.Anything.Value, a3.Anything.Value) {
			t.Fatalf("Values don't match up:\n %+v \n %+v", a, a3)
		}
	}

	if a.Anything.TypeUrl != a2.Anything.TypeUrl {
		if !bytes.Equal(a.Anything.Value, a2.Anything.Value) {
			t.Fatalf("Values don't match up:\n %+v \n %+v", a, a3)
		}
	}

	t.Logf("Values match up:\n %+v \n %+v", a, a3)
	t.Logf("Values match up:\n %+v \n %+v", a, a2)

}
