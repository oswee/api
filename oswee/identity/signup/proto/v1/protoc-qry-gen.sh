#!/bin/sh
# https://gist.github.com/waylan/4080362

ProgName=$(basename $0)
  
sub_help(){
    echo "Usage: $ProgName <subcommand> [options]\n"
    echo "Subcommands:"
    echo "    bar   Do bar"
    echo "    baz   Run baz"
    echo ""
    echo "For help with each subcommand run:"
    echo "$ProgName <subcommand> -h|--help"
    echo ""
}

sub_bar(){
    echo "Running 'bar' command."
}
  
sub_baz(){
    echo "Running 'baz' command."
    echo "First arg is '$1'."
    echo "Second arg is '$2'."
}
  
subcommand=$1
case $subcommand in
    "" | "-h" | "--help")
        sub_help
        ;;
    *)
        shift
        sub_${subcommand} $@
        if [ $? = 127 ]; then
            echo "Error: '$subcommand' is not a known subcommand." >&2
            echo "       Run '$ProgName --help' for a list of known subcommands." >&2
            exit 1
        fi
        ;;
esac

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