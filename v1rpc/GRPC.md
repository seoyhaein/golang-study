# GRPC 다루기전에 먼저 살펴봐야 하는 것들
0. 일단 샘플 파일들은 grpc_clientmain.go grpc_servermain.go 는 패키지 방식을 따르지 않고 그냥 넣어 둠. 향후 옮길 예정임.

1. 기초
   1.1 https://github.com/grpc/grpc-go
   여기서 설치 및 세팅에 대한 이해를 구할 수 있다.
   
   1.2 https://github.com/protocolbuffers/protobuf-go , https://github.com/golang/protobuf
   
   https://pkg.go.dev/google.golang.org/protobuf@v1.27.1
   
   참고) 나는 일반적으로는, import 시  google.golang.org 를 추가해서 사용한다. 예를들어, protobuf file 을 json 으로 encoding/decoding 또는 마샬랑    언마샬링( 의미적으로 차이는 있지만 그냥 쓰도록 함.) 할때 https://github.com/protocolbuffers/protobuf-go 에서 protojson 이 있고                  https://github.com/golang/protobuf 에서 jsonpb 가 있다. 하지만 나는 google.golang.org/protobuf/encoding/protojson 을 사용한다. 정확하진    않지만 해당 github 에서 개발된 녀석들이 공식 패키지로 되는 것이 아닌가 생각해본다. ( 잘못된 정보일 수 있음 )
   
   1.3 https://pkg.go.dev/google.golang.org/protobuf
   protobuf 공식 메뉴얼 웹사이트이다.
   
   1.4 https://developers.google.com/protocol-buffers/docs/overview
   공식 dev guild 사이트 이다.
   
   1.5 https://pkg.go.dev/google.golang.org/grpc#section-documentation
   공식 grpc 메뉴얼이다.

2. grpc 에서 활용되는 패키지들.
   나는 grpc 를 다룰때 server 부분을 golang 으로 작성하고 client 부분은 타 언어로 제작된 패키지가 될 수 있다고 생각한다.
   
   따라서, server 동작시 이를 테스트 해볼 패키지 또는 환경이 마련되어야 한다.
   
   2.1 https://github.com/fullstorydev/grpcurl   
   
   grpcurl 은 간단히 curl 방식으로 grpc server 를 테스트 해볼 수 있는 패키지 이다.
   
   사용방법은 아래와 같다.
   grpcurl 명령어
   proto file 읽기 
   grpcurl --import-path proto 위치 -proto  proto 파일 이름 list 
   
   plaintext insecure server 접속의 경우
   
   grpcurl --plaintext 서버주소(localhost:50051) list > grpc 서비스 호출
   ex> ./grpcurl --plaintext localhost:50052 list
   
   grpcurl --plaintext 서버주소(localhost:50051) list 서비스 이름 > 해당 서비스의 rpc api 호출 
   ex> ./grpcurl --plaintext localhost:50052 list Greeter
   
   grpc 메서드 살펴보기
   grpcurl --plaintext 서버주소(localhost:50051) describe 서비스.메서드 이름
   ex> ./grpcurl --plaintext -d '{"name":"seoyhaein"}' localhost:50052 Greeter.SayHello
   
   grpc 메세지 살펴보기
   grpcurl --plaintext 서버주소(localhost:50051) describe .메세지이름
   
   grpc 메세지 json 형태로 바꾸기
   grpcurl --plaintext -ㅡmgs-template 서버주소(localhost:50051) describe .메세지이름
   
   grpc 메서드 호출
   grpcurl --plaintext -d '{
   "name": "seoyhaein"
   }' 서버주소(localhost:50051) 서비스 이름/메서드 이름
   
   ex> ./grpcurl --plaintext -d '{"name":"seoyhaein"}' localhost:50052 Greeter.SayHello
   
   SayHello 의 입력 파라미터는 HelloRequest 인데 여기서 "name" 이라는 하나의 field 만 있다. 그래서 이렇게 처리한다.
   결과는 아래와 같이, 메세지 파라미터에 구현된 함수에 의해서 "Hello 사용자가 입력한 값" 이 출력되는 형태이다. greet.proto 를 확인하면 금방 이해할 수 있다.
   
   {
  	"message": "Hello seoyhaein"
   }

   
   youtube 에서 설명해주고 있다. hashicorp 에서 개발자로 일하는 사람인거 같은데 (인터뷰에서 몇번 봤음.) 가끔씩 유튜브를 올리는데 내용이 좋다.
   https://www.youtube.com/watch?v=RHWwMrR8LUs
   
   2.2 https://github.com/gogo/protobuf
   
   protobuf 파일을 golang 을 커버팅(컴파일??) 해주는 과정이 필요하다. gogo protobuf를 사용한다. 
   아래와 같이 protoc 에서 gofast 라는 녀석을 주로 사용한다. 물론 다른 것을 사용해도 된다. 이때 생성되는 ~.pb.go 파일은 사용한 컨버팅 툴 (--gofast,etc) 에 따라 다르다.
   생성되는 ~.pb.go 따라 성능에 영향을 미친다고 한다. 세부적으로 비교해보진 않았고 차이만 확인해보았었다.
      
   protoc 를 사용할때 몇가지 문법적?? 인 구문들이 필요한데, 이것들은 추후에 다루도록 한다.
   
   protoc --gofast_out=plugins=grpc:.
   
   2.3 https://github.com/grpc-ecosystem/go-grpc-middleware
   자세히 사용해보진 않았는데, interceptor 부분에서 활용한 적이 있다. 이 패키지는 좀 공부해야할 필요가 있다.
   
   2.4 https://github.com/grpc-ecosystem/go-grpc-prometheus
   프로메테우스 go-grpc 버전인것 같다.
   grpc 모니터링 패키지 이다. 사용해본적이 없지만, 많은 오픈소스에서 활용되고 있는 것을 파악하고 있다.
   
3. https://github.com/grpc/grpc-go/blob/master/server.go 에 대해서

   3.1 NewServer
   
   3.2 RegisterService - reflect 는 차후에 설명, interface 향후 추가 설명 (interface{} 간단히, interface{}{} 도 추후 다뤄보자.)
   
   두개의 파라미터를 사용한다. sd *ServiceDesc 와 ss interface{} 이다. 
   여기서 ss 는 interface{} empty interface 이다. method 를 한개도 가지고 있지 않다. 
   인터페이스에서 가지고 있는 method 를 구현하고 있는 타입? 은 해당 인터페이스를 파라미터로 두고 있는 곳에 위치할 수 있다.
   그런데, empty interface 같은 경우는 method 가 하나도 없어서 그 어떤 type 도 담을 수 없을 것 같지만 반대로 모든 type 을 담을 수 있다.
  
   ```
   // RegisterService registers a service and its implementation to the gRPC
   // server. It is called from the IDL generated code. This must be called before
   // invoking Serve. If ss is non-nil (for legacy code), its type is checked to
   // ensure it implements sd.HandlerType.
   func (s *Server) RegisterService(sd *ServiceDesc, ss interface{}) {
   	if ss != nil {
   		ht := reflect.TypeOf(sd.HandlerType).Elem()
   		st := reflect.TypeOf(ss)
   		if !st.Implements(ht) {
   			logger.Fatalf("grpc: Server.RegisterService found the handler of type %v that does not satisfy %v", st, ht)
   		}
   	}
   	s.register(sd, ss)
   }
   
   // RegisterService 함수의 경우 grpc_servermain.go 에서 사용 예를 살펴보면, protos/greet.pb.go 에서 RegisterGreeterServer 에서 사용하고 있다.
   // srv GreeterServer sms SayHello 메서드를 하나 가지고 있는 인터페이스이다.
  
   func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
   	s.RegisterService(&_Greeter_serviceDesc, srv)
   }
   
   // 여기서 살펴볼것은 아래코드를 보면, greet.pb.go 에서 자동으로 만들어 준 코드인데, grpc_servermain.go 에서 넣어주었다.
   // UnimplementedGreeterServer 같은 경우는 SayHello 를 구현해 주었다. 만약, grpc_servermain.go 에서 SayHello 를 구현해주지 않으면
   // UnimplementedGreeterServer 의 SayHello 가  호출 될 것이다. sample.go 확인.
   
   // UnimplementedGreeterServer can be embedded to have forward compatible implementations.
   type UnimplementedGreeterServer struct {
   }
   
   // status 활용 샘플 추후에 작성하기.
   func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
   	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
   }
   
   // grpc_servermain.go
   type server struct{ pb.UnimplementedGreeterServer }

   // 여기서 _Greeter_serviceDesc 는 protos/greet.pb.go 에서 자동으로 생성되었다. 자세히 보면 greet.proto 파일에서 만들어준 내용들이 들어가 있다.

    var _Greeter_serviceDesc = grpc.ServiceDesc{
	    ServiceName: "Greeter",
	    HandlerType: (*GreeterServer)(nil),
	    Methods: []grpc.MethodDesc{
		    {
			    MethodName: "SayHello",
			    Handler:    _Greeter_SayHello_Handler,
		    },
	    },
	    Streams:  []grpc.StreamDesc{},
	    Metadata: "greet.proto",
    }

   ```
   
   
   
4. 개발 방향
    먼저 기초 샘플을 한번 만들어보고, 좀더 개선된것들을 만들어 본다. 
   기본적인 server 제작 및 interceptor, option 등을 다룬다. 세부적인 것은 추후 논의가 들어갈때 그때 그때 다루도록 한다. 
   https://github.com/grpc/grpc-go/blob/master/server.go
   - NewServer, Server, Serveroptions, RegisterService
   
   tcp/net 세팅에 관련된 부분들도 다루어야 한다. 이건 grpc 랑 상관없지만 결국은 이 녀석들이 들어나게 된다. (이건 좀 시간이 걸리겠군.)
   
   ## grpcurl 에 대한 부분 ( https://github.com/fullstorydev/grpcurl )
   해당 github 주소에서 Makefile 을 보면 release 해주는 부분이 있다.
   
   아래와 같이 최신버전으로 업데이트 해주고, releasing 폴더에서 README.md 파일을 읽어보면 토큰 만들어주고
   (github 토큰 같은 경우는 인터넷 찾아보면 쉽게 만들어주는 방법을 찾을 수 있다.) 
   do-release.sh 실행시켜준 후 make release 하면 release 된 버전을 확인할 수 있다. 나는 그냥 웹사이트에서 다운로드 했다.
   
```
.PHONY: release
release:
	@GO111MODULE=on go install github.com/goreleaser/goreleaser@latest
	goreleaser --rm-dist
```
   
