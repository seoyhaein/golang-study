.PHONY: standard
standard:
	 protoc -I protos/ protos/greet.proto --go_out=plugins=grpc:protos/.

.PHONY: gogofast
gogofast:
	protoc -I protos/ protos/greet.proto --gofast_out=plugins=grpc:protos/gogo/.

.PHONY: explains
explains:
	protoc -I protos/examples/ protos/examples/explains.proto --gofast_out=plugins=grpc:.

# gofast 로 하면 에러

.PHONY: anything
anything:
	protoc -I protos/examples/ protos/examples/anything.proto --go_out=plugins=on

.PHONY: exoneof
exoneof:
	protoc -I protos/examples/ protos/examples/exoneof.proto --go_out=plugins=grpc:protos/examples/.
