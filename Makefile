#DO NOT EDIT! USED FOR DEVELOPMENT

.PHONY: gen_proto

gen_proto:
	@bazel build //api:external_api
	@bazel build //api:internal_api
	@cp bazel-bin/api/external_api_/github.com/YarikRevich/hide-seek-server/internal/api/external-api/v1/proto/external.pb.go internal/api/external-api/v1/proto
	@cp bazel-bin/api/internal_api_/github.com/YarikRevich/hide-seek-server/internal/api/internal-api/v1/proto/internal.pb.go internal/api/internal-api/v1/proto
