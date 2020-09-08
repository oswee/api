# Oswee API

[![lerna](https://img.shields.io/badge/maintained%20with-lerna-cc00ff.svg)](https://lerna.js.org/)
[![lerna](https://img.shields.io/badge/build%20with-bazel-43A047.svg)](https://bazel.build/)

## Important rules

- Every service in separate files
- File names with underscore `file_name.proto`

[Uber guidelines](https://github.com/uber/prototool/blob/dev/style/README.md#directory-structure)

[Google Guidelines](https://cloud.google.com/apis/design/versioning)

Tool to merge swagger files
[go-swagger/go-swagger](https://github.com/go-swagger/go-swagger)

## Prerequisites

Set up GOPATH in `~/.zshrc` or `~/bashrc`

`source $HOME/.zshrc`

```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

In order to compile protobufs, `/usr/local/include` should be populated with [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip).
Copy `include/google/protobuf` directory from downloaded archive into `/usr/local/include/google`

Install grpc-gateway and swagger documentation generator plugins

```sh
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

By default `go get` installs all packages in `$GOPATH`. If that is not specified then in `$HOME/go`.

Copy content of ``%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google`` folder to `/usr/local/include/google` folder.

`cp $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google usr/local/include/google`

Copy `annotations.proto` and `openapiv2.proto` files from `%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options` folder to `/usr/local/include/protoc-gen-swagger/options` folder.

It should look like this:

> :warning: **DEPRECATED**: All external annotations should be included in `./third_party/` directory.

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
└── protoc-gen-swagger
    └── options
        ├── annotations.proto
        └── openapiv2.proto
```

**UPDATE:**
In order to make this repository self-contained, i included vendored packages there and moved Oswee API in its own `/oswee/*` directory.
This way, no matter where this repository will be hosted or used, all Google and Swagger protobufs will be included and i don't need to rely on `/usr/local/include/*`.

**UPDATE2:** It is not an good idea to move `google/protobuf/*` into repository as it is a part of `protoc` compiler and could lead to mismatch between compiler and used definitions.

## Other

Check the version of protoc compiler

```sh
protoc --version
```

Update `protoc` compiler - [StackOverflow](https://stackoverflow.com/a/57776284/6651080)

## Resources

[GitHub Issue about imports not found](https://github.com/grpc-ecosystem/grpc-gateway/issues/574#issuecomment-376018797)

[Protoc Swagger example](https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/proto/examplepb/a_bit_of_everything.proto)

## gRPC endpoint testing

[fullstorydev/grpcui](https://github.com/fullstorydev/grpcui)

## TypeScript

To generate TypeScript types i am using [grpc-web](https://github.com/grpc/grpc-web) plugin.
At time of writing this, TypeScript support is experimental.

### Lerna

Create personal access token

https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token

Create `~/.npmrc` (or at project level) file with Auth token (DO NOT COMMIT AUTH TOKEN!!!)

https://docs.github.com/en/packages/using-github-packages-with-your-projects-ecosystem/configuring-npm-for-use-with-github-packages#authenticating-with-a-personal-access-token

Check the authentication

```sh
npm whoami --registry=https://npm.pkg.github.com
```
