// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.12
// source: tests/import/proto/foo/foo.proto

package foo

import (
	context "context"
	bar "github.com/knqyf263/go-plugin/tests/import/proto/bar"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Request) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Request) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

// go:plugin type=plugin version=1
type Foo interface {
	Hello(context.Context, Request) (bar.Reply, error)
}
