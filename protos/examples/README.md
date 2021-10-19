mesos.proto 는 실제 mesos 에서 사용하는 mesos.proto 를 가져온것이다. 여기서는 proto version 2 이다.
rpc.proto 는 실제 etcd 에서 사용하는 rpc.proto 를 가져온것이다. 여기서는 proto version 3 이다. 
rpc.proto 의경우는 json 호환이다. 자세히 살펴보면 mesos.proto 의 경우는 rpc 가 없다. 즉 mesos.proto 의 경우, http 통신 방법으로 proto 로 통신하는 것을 알 수 있다. 일반적으로 json 으로 http 통신 하는 방식과 그 형태만 다른 것이다.