# GRPC 다루기전에 먼저 살펴봐야 하는 것들
0. 일단 샘플 파일들은 grpc_clientmain.go grpc_servermain.go 는 패키지 방식을 따르지 않고 그냥 넣어 둠. 향후 옮길 예정임.

1. 기초
   1.1 https://github.com/grpc/grpc-go
   여기서 설치 및 세팅에 대한 이해를 구할 수 있다.
   
   1.2 https://github.com/protocolbuffers/protobuf-go
   
   1.3 https://pkg.go.dev/google.golang.org/protobuf
   protobuf 공식 메뉴얼 웹사이트이다.
   
   1.4 https://developers.google.com/protocol-buffers/docs/overview
   공식 dev guild 사이트 이다.

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
   
3. 개발 방향
    먼저 기초 샘플을 한번 만들어보고, 좀더 개선된것들을 만들어 본다. 
   기본적인 server 제작 및 interceptor, option 등을 다룬다. 세부적인 것은 추후 논의가 들어갈때 그때 그때 다루도록 한다. 
   https://github.com/grpc/grpc-go/blob/master/server.go
   - NewServer, Server, Serveroptions, RegisterService
   
   tcp/net 세팅에 관련된 부분들도 다루어야 한다. 이건 grpc 랑 상관없지만 결국은 이 녀석들이 들어나게 된다. (이건 좀 시간이 걸리겠군.)
   
   ## grpcurl 에 대한 부분
   Makefile 을 보면 release 해주는 부분이 있다.
   
   아래와 같이 최신버전으로 업데이트 해주고, releasing 폴더에서 README.md 파일을 읽어보면 토큰 만들어주고
   (github 토큰 같은 경우는 인터넷 찾아보면 쉽게 만들어주는 방법을 찾을 수 있다.) 
   do-release.sh 실행시켜준 후 make release 하면 release 된 버전을 확인할 수 있다. 나는 그냥 웹사이트에서 다운로드 했다.
   
```
.PHONY: release
release:
	@GO111MODULE=on go install github.com/goreleaser/goreleaser@latest
	goreleaser --rm-dist
```
   
