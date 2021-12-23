load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/YarikRevich/HideSeek-Server
gazelle(name = "gazelle")

go_library(
    name = "HideSeek-Server_lib",
    srcs = ["main.go"],
    importpath = "github.com/YarikRevich/HideSeek-Server",
    visibility = ["//visibility:private"],
    deps = [
        "//internal/api/external-api/v1/implementation",
        "//internal/api/external-api/v1/proto",
        "//internal/cache",
        "//internal/interceptors",
        "//internal/monitoring",
        "//tools/params",
        "//tools/printer",
        "@com_github_sirupsen_logrus//:go_default_library",
        "@com_github_yarikrevich_go_demonizer//pkg/demonizer:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//encoding/gzip:go_default_library",
    ],
)

go_binary(
    name = "HideSeek-Server",
    embed = [":HideSeek-Server_lib"],
    visibility = ["//visibility:public"],
)
