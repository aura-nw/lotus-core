#!/usr/bin/env bash

# Find the go bin directory and add it to the PATH.
gobin=$(go env GOBIN)
gopath=$(go env GOPATH)
if test -z $gobin; then
  if test -z $gopath; then
    gopath="$HOME/go"
  fi
  gobin="$gopath/bin"
fi
export PATH="$PATH:$gobin"

# Check that needed binaries are available.
protoc=$(which protoc)
if test -z $protoc; then
  printf "protoc binary not found.  Please run:\n\tsudo apt install protobuf-compiler\n, and then re-run this command.\n"
  exit 1
fi
gengo=$(which protoc-gen-go)
if test -z $gengo; then
  printf "protoc-gen-go binary not found.  Please run:\n\tgo install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26\nand then re-run this command."
  exit 1
fi

exec protoc --go_out=. --go_opt=paths=source_relative,Mgoogle/protobuf/timestamp.proto=google.golang.org/protobuf/types/known/timestamppb ${1+"$@"}