#!/bin/sh

protoc \
    --proto_path=signup/proto/v1 \
    --go_out=plugins=grpc:signup/stubs/v1 \
    query-service.proto

protoc \
    --proto_path=signup/proto/v1 \
    --grpc-gateway_out=logtostderr=true:signup/stubs/v1 \
    query-service.proto

protoc \
    --proto_path=signup/proto/v1 \
    --swagger_out=logtostderr=true:signup/swagger/v1 \
    query-service.proto