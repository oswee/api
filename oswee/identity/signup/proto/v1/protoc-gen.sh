#!/bin/sh

# Command API

protoc \
    --proto_path=. \
    --go_out=plugins=grpc:. \
    oswee/identity/signup/proto/v1/signup.proto

protoc \
    --proto_path=. \
    --grpc-gateway_out=logtostderr=true:. \
    oswee/identity/signup/proto/v1/signup.proto

# Command API

protoc \
    --proto_path=. \
    --go_out=plugins=grpc:. \
    oswee/identity/signup/proto/v1/signup_command_api.proto

protoc \
    --proto_path=. \
    --grpc-gateway_out=logtostderr=true:. \
    oswee/identity/signup/proto/v1/signup_command_api.proto

protoc \
    --proto_path=. \
    --proto_path=oswee/identity/signup/proto/v1 \
    --swagger_out=logtostderr=true:oswee/identity/signup/swagger/v1 \
    signup_command_api.proto

# Query API

protoc \
    --proto_path=. \
    --go_out=plugins=grpc:. \
    oswee/identity/signup/proto/v1/signup_query_api.proto

protoc \
    --proto_path=. \
    --grpc-gateway_out=logtostderr=true:. \
    oswee/identity/signup/proto/v1/signup_query_api.proto

protoc \
    --proto_path=. \
    --proto_path=oswee/identity/signup/proto/v1 \
    --swagger_out=logtostderr=true:oswee/identity/signup/swagger/v1 \
    signup_query_api.proto

# Command Handler

protoc \
    --proto_path=. \
    --go_out=plugins=grpc:. \
    oswee/identity/signup/proto/v1/signup_command_handler.proto