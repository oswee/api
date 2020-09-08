load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_typescript_proto//:index.bzl", "typescript_proto_library")

proto_library(
    name = "module_proto",
    srcs = ["module.proto"],
    strip_import_prefix = "/oswee",
    visibility = ["//visibility:public"],
    deps = ["@go_googleapis//google/api:annotations_proto"],
)

go_proto_library(
    name = "module_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/oswee/api/oswee/module/stubs/v1",
    proto = ":module_proto",
    visibility = ["//visibility:public"],
    deps = ["@go_googleapis//google/api:annotations_go_proto"],
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/oswee/api/oswee/module/proto/v1",
    visibility = ["//visibility:public"],
)

typescript_proto_library(
    name = "module_ts_proto",
    proto = ":module_proto",
)
