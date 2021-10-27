## protobuf 에 대해서
mesos.proto 는 실제 mesos 에서 사용하는 mesos.proto 를 가져온것이다. 여기서는 proto version 2 이다.
rpc.proto 는 실제 etcd 에서 사용하는 rpc.proto 를 가져온것이다. 여기서는 proto version 3 이다. 
rpc.proto 의경우는 json 호환이다. 

자세히 살펴보면 mesos.proto 의 경우는 service (rpc) 가 없다. 즉 mesos.proto 의 경우, http 통신 방법으로 proto 로 통신하는 것을 알 수 있다. 
일반적으로 json 으로 http 통신 하는 방식과 그 형태만 다른 것이다.

rpc.proto 의 경우는 grpc 를 통해서 통신하는 것을 확인 할 수 있다. 즉, 같은 proto 를 쓰고 있지만 하나는 http 방식, 하나는 grpc 방식으로 통신하는 것을 알 수 있다.

## 정리 순서 (소스 수정해야함. 일단 막 작성중)

https://developers.google.com/protocol-buffers/docs/proto3#maps

를 기반으로 정리중. 일단 빠르게 작성하고 코드 및 문서에 대한 내용 정리는 다 마친 후에 하자.

1. protobuf 정리중 상태
   1.6 map
   ...
   1.X techniques
   
2. gogoprotobuf 


참고 : https://jeong-pro.tistory.com/193
example> https://github.com/seoyhaein/conductor/blob/main/grpc/src/main/proto/model/task.proto
   