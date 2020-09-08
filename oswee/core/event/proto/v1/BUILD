load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "v1_proto",
    srcs = [
        "event.proto",
        "service.proto",
    ],
    strip_import_prefix = "/core/event/proto",
    visibility = ["//visibility:public"],
    deps = [
        "//:root_proto",
        "//protoc-gen-swagger/options:options_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "oswee/core/event/stubs/v1",
    proto = ":v1_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//:go_default_library",
        "//protoc-gen-swagger/options:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":v1_go_proto"],
    importpath = "oswee/core/event/stubs/v1",
    visibility = ["//visibility:public"],
)
