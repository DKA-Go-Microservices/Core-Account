//go:build tools

package tools

import (
	_ "github.com/favadi/protoc-go-inject-tag"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
