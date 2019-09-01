#!/bin/sh

protoc \
    --proto_path=signup/proto/v1 \
    --proto_path=vendor \
    --go_out=plugins=grpc:signup/stubs/v1 \
    command-service.proto

protoc \
    --proto_path=signup/proto/v1 \
    --proto_path=vendor \
    --grpc-gateway_out=logtostderr=true:signup/stubs/v1 \
    command-service.proto

protoc \
    --proto_path=signup/proto/v1 \
    --proto_path=vendor \
    --swagger_out=logtostderr=true:signup/swagger/v1 \
    command-service.proto