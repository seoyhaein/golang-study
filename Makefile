.PHONY: standard
standard:
	 protoc -I protos/ protos/greet.proto --go_out=plugins=grpc:protos/.

.PHONY: gogofast
gogofast:
	protoc -I protos/ protos/greet.proto --gofast_out=plugins=grpc:protos/gogo/.

.PHONY: explains
explains:
	protoc -I protos protos/examples/explains.proto --gofast_out=plugins=grpc:.
