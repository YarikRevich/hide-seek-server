.PHONY: all build install

all: gen_api build install 

API_VERSION := v1
API_TYPE := external

gen_api:
	@protoc -I internal/api/$(API_TYPE)-api/$(API_VERSION)/proto --go_out=. api.proto
	@protoc -I internal/api/$(API_TYPE)-api/$(API_VERSION)/proto --go-grpc_out=. api.proto

test:
	@go test ./... 

build:
	@go build

install:
	@go install
