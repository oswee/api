load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "examplepb_proto",
    srcs = ["example.proto"],
    strip_import_prefix = "/example/proto",
    visibility = ["//visibility:public"],
    deps = [
        "//examples/proto/pathenum:pathenum_proto",
        "//examples/proto/sub:sub_proto",
        "//examples/proto/sub2:sub2_proto",
        "//protoc-gen-swagger/options:options_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:field_mask_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "examplepb_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/oswee/api/oswee/example/proto/v1",
    proto = ":examplepb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//examples/proto/pathenum:go_default_library",
        "//examples/proto/sub:go_default_library",
        "//examples/proto/sub2:go_default_library",
        "//protoc-gen-swagger/options:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    embed = [":examplepb_go_proto"],
    importpath = "github.com/oswee/api/oswee/example/proto/v1",
    visibility = ["//visibility:public"],
)
