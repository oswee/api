#!/bin/sh

# Command API

protoc \
    --proto_path=. \
    --go_out=plugins=grpc:. \
    oswee/core/event/proto/v1/event.proto

protoc \
    --proto_path=. \
    --grpc-gateway_out=logtostderr=true:. \
    oswee/core/event/proto/v1/event.proto


