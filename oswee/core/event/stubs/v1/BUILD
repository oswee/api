load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["event.pb.go"],
    importpath = "github.com/oswee/api/oswee/core/event/stubs/v1",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@io_bazel_rules_go//proto/wkt:timestamp_go_proto",
    ],
)
