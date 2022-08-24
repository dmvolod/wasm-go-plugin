// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.0.1
// 	protoc               v3.21.5
// source: test/fields/proto/fields.proto

package proto

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_A Enum = 0
	Enum_B Enum = 1
	Enum_C Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	Enum_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float64              `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float32              `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
	C int32                `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D int64                `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E uint32               `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
	F uint64               `protobuf:"varint,6,opt,name=f,proto3" json:"f,omitempty"`
	G int32                `protobuf:"zigzag32,7,opt,name=g,proto3" json:"g,omitempty"`
	H int64                `protobuf:"zigzag64,8,opt,name=h,proto3" json:"h,omitempty"`
	I uint32               `protobuf:"fixed32,9,opt,name=i,proto3" json:"i,omitempty"`
	J uint64               `protobuf:"fixed64,10,opt,name=j,proto3" json:"j,omitempty"`
	K int32                `protobuf:"fixed32,11,opt,name=k,proto3" json:"k,omitempty"`
	L int64                `protobuf:"fixed64,12,opt,name=l,proto3" json:"l,omitempty"`
	M bool                 `protobuf:"varint,13,opt,name=m,proto3" json:"m,omitempty"`
	N string               `protobuf:"bytes,14,opt,name=n,proto3" json:"n,omitempty"`
	O []byte               `protobuf:"bytes,15,opt,name=o,proto3" json:"o,omitempty"`
	P []string             `protobuf:"bytes,16,rep,name=p,proto3" json:"p,omitempty"`
	Q map[string]*IntValue `protobuf:"bytes,17,rep,name=q,proto3" json:"q,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	R *Request_Nested      `protobuf:"bytes,18,opt,name=r,proto3" json:"r,omitempty"`
	S Enum                 `protobuf:"varint,19,opt,name=s,proto3,enum=greeting.Enum" json:"s,omitempty"`
}

func (x *Request) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Request) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Request) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Request) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Request) GetD() int64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *Request) GetE() uint32 {
	if x != nil {
		return x.E
	}
	return 0
}

func (x *Request) GetF() uint64 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *Request) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *Request) GetH() int64 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *Request) GetI() uint32 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *Request) GetJ() uint64 {
	if x != nil {
		return x.J
	}
	return 0
}

func (x *Request) GetK() int32 {
	if x != nil {
		return x.K
	}
	return 0
}

func (x *Request) GetL() int64 {
	if x != nil {
		return x.L
	}
	return 0
}

func (x *Request) GetM() bool {
	if x != nil {
		return x.M
	}
	return false
}

func (x *Request) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *Request) GetO() []byte {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *Request) GetP() []string {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *Request) GetQ() map[string]*IntValue {
	if x != nil {
		return x.Q
	}
	return nil
}

func (x *Request) GetR() *Request_Nested {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Request) GetS() Enum {
	if x != nil {
		return x.S
	}
	return Enum_A
}

type IntValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *IntValue) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *IntValue) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float64              `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float32              `protobuf:"fixed32,2,opt,name=b,proto3" json:"b,omitempty"`
	C int32                `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D int64                `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E uint32               `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
	F uint64               `protobuf:"varint,6,opt,name=f,proto3" json:"f,omitempty"`
	G int32                `protobuf:"zigzag32,7,opt,name=g,proto3" json:"g,omitempty"`
	H int64                `protobuf:"zigzag64,8,opt,name=h,proto3" json:"h,omitempty"`
	I uint32               `protobuf:"fixed32,9,opt,name=i,proto3" json:"i,omitempty"`
	J uint64               `protobuf:"fixed64,10,opt,name=j,proto3" json:"j,omitempty"`
	K int32                `protobuf:"fixed32,11,opt,name=k,proto3" json:"k,omitempty"`
	L int64                `protobuf:"fixed64,12,opt,name=l,proto3" json:"l,omitempty"`
	M bool                 `protobuf:"varint,13,opt,name=m,proto3" json:"m,omitempty"`
	N string               `protobuf:"bytes,14,opt,name=n,proto3" json:"n,omitempty"`
	O []byte               `protobuf:"bytes,15,opt,name=o,proto3" json:"o,omitempty"`
	P []string             `protobuf:"bytes,16,rep,name=p,proto3" json:"p,omitempty"`
	Q map[string]*IntValue `protobuf:"bytes,17,rep,name=q,proto3" json:"q,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	R *Response_Nested     `protobuf:"bytes,18,opt,name=r,proto3" json:"r,omitempty"`
	S Enum                 `protobuf:"varint,19,opt,name=s,proto3,enum=greeting.Enum" json:"s,omitempty"`
}

func (x *Response) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Response) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Response) GetB() float32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Response) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Response) GetD() int64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *Response) GetE() uint32 {
	if x != nil {
		return x.E
	}
	return 0
}

func (x *Response) GetF() uint64 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *Response) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *Response) GetH() int64 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *Response) GetI() uint32 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *Response) GetJ() uint64 {
	if x != nil {
		return x.J
	}
	return 0
}

func (x *Response) GetK() int32 {
	if x != nil {
		return x.K
	}
	return 0
}

func (x *Response) GetL() int64 {
	if x != nil {
		return x.L
	}
	return 0
}

func (x *Response) GetM() bool {
	if x != nil {
		return x.M
	}
	return false
}

func (x *Response) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *Response) GetO() []byte {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *Response) GetP() []string {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *Response) GetQ() map[string]*IntValue {
	if x != nil {
		return x.Q
	}
	return nil
}

func (x *Response) GetR() *Response_Nested {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Response) GetS() Enum {
	if x != nil {
		return x.S
	}
	return Enum_A
}

type Request_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Request_Nested) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Request_Nested) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

type Response_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *Response_Nested) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Response_Nested) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

// go:plugin type=plugin version=1
type FieldTest interface {
	Test(context.Context, Request) (Response, error)
}
type EmptyHostFunctions interface {
}
