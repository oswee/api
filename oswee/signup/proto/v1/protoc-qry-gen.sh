#!/bin/sh

protoc \
    --proto_path=. \
    --go_out=plugins=grpc:. \
    oswee/signup/proto/v1/query-service.proto

protoc \
    --proto_path=. \
    --grpc-gateway_out=logtostderr=true:. \
    oswee/signup/proto/v1/query-service.proto

protoc \
    --proto_path=. \
    --proto_path=oswee/signup/proto/v1 \
    --swagger_out=logtostderr=true:oswee/signup/swagger/v1 \
    query-service.proto