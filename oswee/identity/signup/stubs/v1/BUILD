load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "signup.pb.go",
        "signup_command_api.pb.go",
        "signup_command_api.pb.gw.go",
        "signup_command_handler.pb.go",
        "signup_event_handler.pb.go",
        "signup_projection.pb.go",
        "signup_query_api.pb.go",
        "signup_query_api.pb.gw.go",
    ],
    importpath = "github.com/oswee/api/oswee/identity/signup/stubs/v1",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_grpc_ecosystem_grpc_gateway//protoc-gen-swagger/options:go_default_library",
        "@com_github_grpc_ecosystem_grpc_gateway//runtime:go_default_library",
        "@com_github_grpc_ecosystem_grpc_gateway//utilities:go_default_library",
        "@go_googleapis//google/api:annotations_go_proto",
        "@io_bazel_rules_go//proto/wkt:empty_go_proto",
        "@io_bazel_rules_go//proto/wkt:timestamp_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//grpclog:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
    ],
)
