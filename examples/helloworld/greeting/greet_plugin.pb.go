//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.12
// source: examples/helloworld/greeting/greet.proto

package greeting

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

const GreeterPluginAPIVersion = 1

//export greeter_api_version
func _greeter_api_version() uint64 {
	return GreeterPluginAPIVersion
}

var greeter Greeter

func RegisterGreeter(p Greeter) {
	greeter = p
}

//export greeter_greet
func _greeter_greet(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req GreetRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := greeter.Greet(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
