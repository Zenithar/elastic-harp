// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: cso/v1/validator_api.proto

package csov1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ValidatorAPIClient is the client API for ValidatorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ValidatorAPIClient interface {
	// Validate given path according to CSO sepcification.
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type validatorAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewValidatorAPIClient(cc grpc.ClientConnInterface) ValidatorAPIClient {
	return &validatorAPIClient{cc}
}

func (c *validatorAPIClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/cso.v1.ValidatorAPI/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ValidatorAPIServer is the server API for ValidatorAPI service.
// All implementations should embed UnimplementedValidatorAPIServer
// for forward compatibility
type ValidatorAPIServer interface {
	// Validate given path according to CSO sepcification.
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

// UnimplementedValidatorAPIServer should be embedded to have forward compatible implementations.
type UnimplementedValidatorAPIServer struct {
}

func (UnimplementedValidatorAPIServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

// UnsafeValidatorAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ValidatorAPIServer will
// result in compilation errors.
type UnsafeValidatorAPIServer interface {
	mustEmbedUnimplementedValidatorAPIServer()
}

func RegisterValidatorAPIServer(s grpc.ServiceRegistrar, srv ValidatorAPIServer) {
	s.RegisterService(&ValidatorAPI_ServiceDesc, srv)
}

func _ValidatorAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ValidatorAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cso.v1.ValidatorAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ValidatorAPIServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ValidatorAPI_ServiceDesc is the grpc.ServiceDesc for ValidatorAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ValidatorAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cso.v1.ValidatorAPI",
	HandlerType: (*ValidatorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _ValidatorAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cso/v1/validator_api.proto",
}
