.PHONY: standard
standard:
	 protoc -I protos/ protos/greet.proto --go_out=plugins=grpc:protos/.

.PHONY: gogofast
gogofast:
	protoc -I protos/ protos/greet.proto --gofast_out=plugins=grpc:protos/gogo/.


# gofast 로 하면 에러

.PHONY: exany
exany:
	protoc -I protos/examples/ protos/examples/exany.proto --go_out=plugins=grpc:protos/examples/.

.PHONY: exoneof
exoneof:
	protoc -I protos/examples/ protos/examples/exoneof.proto --go_out=plugins=grpc:protos/examples/.

.PHONY: exscala
exscala:
	protoc -I protos/examples/ protos/examples/exscala.proto --go_out=plugins=grpc:protos/examples/.

.PHONY: all
all:
	protoc -I protos/examples/*.proto --go_out=plugins=grpc:protos/examples/.

.PHONY: tasks
tasks:
	protoc -I protos/ protos/tasks.proto --gofast_out=plugins=grpc:protos/.

## genprotoc.sh 참고

