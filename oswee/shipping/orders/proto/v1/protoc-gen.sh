#!/bin/sh

# Command API

protoc \
    --proto_path=third_party \
    --proto_path=oswee/shipping/orders/proto/v1 \
    --go_out=plugins=grpc:. \
    oswee/shipping/orders/proto/v1/order.proto

protoc \
    --proto_path=third_party \
    --proto_path=oswee/shipping/orders/proto/v1 \
    --grpc-gateway_out=logtostderr=true:. \
    oswee/shipping/orders/proto/v1/order.proto

# Command API

protoc \
  -I. \
  --proto_path=third_party \
  --go_out=plugins=grpc:. \
  oswee/shipping/orders/proto/v1/orders_command_api.proto

# protoc \
#     --proto_path=third_party \
#     --proto_path=oswee/shipping/orders/proto/v1 \
#     --grpc-gateway_out=logtostderr=true:. \
#     oswee/shipping/orders/proto/v1/orders_command_api.proto

# protoc \
#     --proto_path=third_party \
#     --proto_path=oswee/shipping/orders/proto/v1 \
#     --swagger_out=logtostderr=true:oswee/shipping/orders/swagger/v1 \
#     orders_command_api.proto
