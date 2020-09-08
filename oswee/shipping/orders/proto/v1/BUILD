load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "v1_proto",
    srcs = [
        "order.proto",
        "orders_command_api.proto",
    ],
    strip_import_prefix = "/shipping/orders/proto",
    visibility = ["//visibility:public"],
    deps = [
        "//oswee/shipping/orders/proto/v1:v1_proto",
        "//protoc-gen-swagger/options:options_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "v1_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "oswee/shipping/orders/stubs/v1",
    proto = ":v1_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//oswee/shipping/orders/proto/v1:go_default_library",
        "//protoc-gen-swagger/options:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":v1_go_proto"],
    importpath = "oswee/shipping/orders/stubs/v1",
    visibility = ["//visibility:public"],
)
