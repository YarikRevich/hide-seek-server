.PHONY: help test gen_proto prepare install_deps build run

#System vars
OS := $(shell uname -s)
SHELL := /bin/bash

test:
	@bazel query //... | xargs bazel test;

gen_proto:
	
	@cp bazel-bin/api/external_proto_go_/internal/api/external-api/v1/proto/external.pb.go internal/api/external-api/v1/proto
	@cp bazel-bin/api/internal_proto_go_/internal/api/internal-api/v1/proto/internal.pb.go internal/api/internal-api/v1/proto

run:
# @cp bazel-bin/cmd/hide-seek-server_/hide-seek-server /usr/local/bin;

# ifeq ($(type), service)
# ifeq ($(OS), Linux)
# 	@chown root:wheel configs/hide-seek-server.service
# 	@chmod 600 configs/hide-seek-server.service
# 	@cp configs/hide-seek-server.service /etc/systemd/system
# 	@systemctl daemon-reload
# 	@systemctl start hide-seek-server.service
# 	@systemctl enable hide-seek-server.service
# else ifeq ($(OS), Darwin)
# 	@chown root:wheel configs/hide-seek-server.plist
# 	@chmod 600 configs/hide-seek-server.plist
# 	@launchctl load configs/hide-seek-server.plist
# 	@launchctl start hide-seek-server
# endif
# endif
	@docker network create hide-seek-server
	@docker build -t hide-seek -f configs/Dockerfile.server .
	@docker run -d --restart=always --network=hide-seek --name=hide-seek-server hide-seek-server

	@bazel test //configs:test_prometheus_config;
	@docker build -t prometheus -f configs/Dockerfile.prometheus .
	@docker run -d --restart=always --network=hide-seek --name=prometheus prometheus

	@bazel run //configs:grafana_docker_image;
	@docker rm -f --restart=always $$(docker ps -q --filter="ancestor=bazel/configs:grafana_docker_image");
	@docker run -p 3000:3000 -d --restart=always --network=hide-seek --name=grafana bazel/configs:grafana_docker_image;