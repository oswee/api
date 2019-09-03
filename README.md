# Oswee API

### Prerequisites

In order to compile protobufs, `/usr/local/include` should be populated with [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip).
Copy `include/google/protobuf` directory from downloaded archive into `/usr/local/include/google`


Install grpc-gateway and swagger documentation generator plugins

```sh
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

Copy content of `%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google` folder to `/usr/local/include/google` folder.

Copy `annotations.proto` and `openapiv2.proto` files from `%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options` folder to `/usr/local/include/protoc-gen-swagger/options` folder.

It should look like this:

```sh
/usr/local/include
├── google
│   ├── api
│   │   ├── annotations.proto
│   │   └── http.proto
│   ├── protobuf
│   │   ├── any.proto
│   │   ├── api.proto
│   │   ├── compiler
│   │   │   └── plugin.proto
│   │   ├── descriptor.proto
│   │   ├── duration.proto
│   │   ├── empty.proto
│   │   ├── field_mask.proto
│   │   ├── source_context.proto
│   │   ├── struct.proto
│   │   ├── timestamp.proto
│   │   ├── type.proto
│   │   └── wrappers.proto
│   └── rpc
│       ├── code.proto
│       ├── error_details.proto
│       └── status.proto
├── protoc-gen-swagger
│   └── options
│       ├── annotations.proto
│       └── openapiv2.proto
```

### Other

Check the version of protoc compiler

```sh
protoc --version
```

Update `protoc` compiler - [StackOverflow](https://stackoverflow.com/a/57776284/6651080)

### Resources

[GitHub Issue about imports not found](https://github.com/grpc-ecosystem/grpc-gateway/issues/574#issuecomment-376018797)

[Protoc Swagger example](https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/proto/examplepb/a_bit_of_everything.proto)