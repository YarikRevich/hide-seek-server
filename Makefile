.PHONY: default help test build install

default: help

help: 
	@echo "It works"

gen_proto:
	@protoc -I api --go_out=. external.proto internal.proto
	@protoc -I api --go-grpc_out=. external.proto internal.proto

test:
	@go test ./... 

build:
	@go build -o HideSeek-Server cmd/main.go

install:
	@go install
