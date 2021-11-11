.PHONY: all build install

all: build install

gen_proto:
	@protoc -I internal/api --go_out=. api.proto
	@protoc -I internal/api --go-grpc_out=. api.proto

build:
	@go build

install:
	@go install
