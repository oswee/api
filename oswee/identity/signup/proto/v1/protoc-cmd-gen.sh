#!/bin/sh

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