.PHONY: help test gen_proto build install

OS := $(shell uname -s)
installation_type ?= executable

help:
	@echo "These are all available commands:\n\n--build: builds all the components('hide-seek-server', 'hide-seek-client', 'hide-seek-services')\n--install: installs all the components('hide-seek-server', 'hide-seek-client', 'hide-seek-services')\n";

test:
	@bazel test ./...

gen_proto:
	@bazel build //api:internal_proto_go
	@bazel build //api:external_proto_go
	
	@cp bazel-bin/api/external_proto_go_/internal/api/external-api/v1/proto/external.pb.go internal/api/external-api/v1/proto
	@cp bazel-bin/api/internal_proto_go_/internal/api/internal-api/v1/proto/internal.pb.go internal/api/internal-api/v1/proto

	# @protoc -I api --go_out=. external.proto internal.proto
	# @protoc -I api --go-grpc_out=. external.proto internal.proto

build: gen_proto
	@bazel build //cmd:hide-seek-server

install:
	@cp bazel-bin/cmd/hide-seek-server_/hide-seek-server /usr/local/bin;

ifeq ($(installation_type), service)
ifeq ($(OS), Linux)
	@chown root:wheel configs/hide-seek-server.service
	@chmod 600 configs/hide-seek-server.service
	@cp configs/hide-seek-server.service /etc/systemd/system
	@systemctl daemon-reload
	@systemctl start hide-seek-server.service
	@systemctl enable hide-seek-server.service
else ifeq ($(OS), Darwin)
	@chown root:wheel configs/hide-seek-server.plist
	@chmod 600 configs/hide-seek-server.plist
	@launchctl load configs/hide-seek-server.plist
	@launchctl start hide-seek-server
endif
endif