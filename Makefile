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

# https://developers.google.com/protocol-buffers/docs/reference/go-generated
# import *.proto 같은 경우는 위의 링크를 참고해서 확실하게 하자.define

#GOPATH 설정이 안되었을 경우. 현재 사용하는 컴은 지금은 설정이 되어 있다.
# 참고 :  https://developers.google.com/protocol-buffers/docs/reference/go-generated

# detect GOPATH if not set
#ifndef $(GOPATH)
#    $(info GOPATH is not set, autodetecting..)
#    TESTPATH := $(dir $(abspath ../../..))
#    DIRS := bin pkg src
    # create a ; separated line of tests and pass it to shell
#    MISSING_DIRS := $(shell $(foreach entry,$(DIRS),test -d "$(TESTPATH)$(entry)" || echo "$(entry)";))
#    ifeq ($(MISSING_DIRS),)
#        $(info Found GOPATH: $(TESTPATH))
#        export GOPATH := $(TESTPATH)
#    else
#        $(info ..missing dirs "$(MISSING_DIRS)" in "$(TESTDIR)")
#        $(info GOPATH autodetection failed)
#    endif
#endif

.PHONY: jobs
jobs:
	protoc --proto_path=$(GOPATH)/src:.	\
			./protos/jobs.proto \
		   --go_out=plugins=grpc:.

## genprotoc.sh 참고

