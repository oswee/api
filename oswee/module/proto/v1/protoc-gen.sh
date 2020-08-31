#!/bin/sh

# Run from rpository root as ./oswee/module/proto/v1/protoc-gen.sh

protoc \
    --proto_path=third_party \
    --proto_path=oswee/module/proto/v1 \
    --go_out=plugins=grpc:. \
    --js_out=import_style=commonjs:oswee/module/stubs/v1 \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:oswee/module/stubs/v1 \
    oswee/module/proto/v1/module.proto
