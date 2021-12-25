load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/YarikRevich/hide-seek-server
gazelle(name = "gazelle")

go_library(
    name = "hide-seek-server_lib",
    srcs = ["main.go"],
    importpath = "github.com/YarikRevich/hide-seek-server",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_yarikrevich_go_demonizer//pkg/demonizer:go_default_library",
        "@com_github_yarikrevich_hideseek_server//tools/printer:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//encoding/gzip:go_default_library",
    ],
)

go_binary(
    name = "hide-seek-server",
    embed = [":hide-seek-server_lib"],
    visibility = ["//visibility:public"],
)
