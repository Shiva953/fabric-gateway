//
//Copyright 2020 IBM All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// compile using
// protoc -I. -I$GOPATH/src/github.com/hyperledger/fabric-protos --go_out=plugins=grpc:. --go_opt=paths=source_relative protos/gateway.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.3
// source: protos/gateway.proto

package gateway

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	common "github.com/hyperledger/fabric-protos-go/common"
	peer "github.com/hyperledger/fabric-protos-go/peer"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The value that is returned by the transaction function
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_protos_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// The signed proposal ready for endorsement plus any processing options
type ProposedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proposal *peer.SignedProposal `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *ProposedTransaction) Reset() {
	*x = ProposedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposedTransaction) ProtoMessage() {}

func (x *ProposedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposedTransaction.ProtoReflect.Descriptor instead.
func (*ProposedTransaction) Descriptor() ([]byte, []int) {
	return file_protos_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *ProposedTransaction) GetProposal() *peer.SignedProposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

// The set of transaction responses from the endorsing peers for signing by the client
// before submitting to ordering service (via gateway)
type PreparedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId     string           `protobuf:"bytes,1,opt,name=txId,proto3" json:"txId,omitempty"`
	Response *Result          `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	Envelope *common.Envelope `protobuf:"bytes,3,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *PreparedTransaction) Reset() {
	*x = PreparedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparedTransaction) ProtoMessage() {}

func (x *PreparedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparedTransaction.ProtoReflect.Descriptor instead.
func (*PreparedTransaction) Descriptor() ([]byte, []int) {
	return file_protos_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *PreparedTransaction) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *PreparedTransaction) GetResponse() *Result {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *PreparedTransaction) GetEnvelope() *common.Envelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_protos_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_protos_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_protos_gateway_proto protoreflect.FileDescriptor

var file_protos_gateway_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x13,
	0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x78, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52,
	0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xc5, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x45, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x08, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x67, 0x0a, 0x26, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x0c, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_gateway_proto_rawDescOnce sync.Once
	file_protos_gateway_proto_rawDescData = file_protos_gateway_proto_rawDesc
)

func file_protos_gateway_proto_rawDescGZIP() []byte {
	file_protos_gateway_proto_rawDescOnce.Do(func() {
		file_protos_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_gateway_proto_rawDescData)
	})
	return file_protos_gateway_proto_rawDescData
}

var file_protos_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_gateway_proto_goTypes = []interface{}{
	(*Result)(nil),              // 0: protos.Result
	(*ProposedTransaction)(nil), // 1: protos.ProposedTransaction
	(*PreparedTransaction)(nil), // 2: protos.PreparedTransaction
	(*Event)(nil),               // 3: protos.Event
	(*peer.SignedProposal)(nil), // 4: protos.SignedProposal
	(*common.Envelope)(nil),     // 5: common.Envelope
}
var file_protos_gateway_proto_depIdxs = []int32{
	4, // 0: protos.ProposedTransaction.proposal:type_name -> protos.SignedProposal
	0, // 1: protos.PreparedTransaction.response:type_name -> protos.Result
	5, // 2: protos.PreparedTransaction.envelope:type_name -> common.Envelope
	1, // 3: protos.Gateway.Prepare:input_type -> protos.ProposedTransaction
	2, // 4: protos.Gateway.Commit:input_type -> protos.PreparedTransaction
	1, // 5: protos.Gateway.Evaluate:input_type -> protos.ProposedTransaction
	2, // 6: protos.Gateway.Prepare:output_type -> protos.PreparedTransaction
	3, // 7: protos.Gateway.Commit:output_type -> protos.Event
	0, // 8: protos.Gateway.Evaluate:output_type -> protos.Result
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_gateway_proto_init() }
func file_protos_gateway_proto_init() {
	if File_protos_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposedTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparedTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_gateway_proto_goTypes,
		DependencyIndexes: file_protos_gateway_proto_depIdxs,
		MessageInfos:      file_protos_gateway_proto_msgTypes,
	}.Build()
	File_protos_gateway_proto = out.File
	file_protos_gateway_proto_rawDesc = nil
	file_protos_gateway_proto_goTypes = nil
	file_protos_gateway_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	Prepare(ctx context.Context, in *ProposedTransaction, opts ...grpc.CallOption) (*PreparedTransaction, error)
	Commit(ctx context.Context, in *PreparedTransaction, opts ...grpc.CallOption) (Gateway_CommitClient, error)
	Evaluate(ctx context.Context, in *ProposedTransaction, opts ...grpc.CallOption) (*Result, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Prepare(ctx context.Context, in *ProposedTransaction, opts ...grpc.CallOption) (*PreparedTransaction, error) {
	out := new(PreparedTransaction)
	err := c.cc.Invoke(ctx, "/protos.Gateway/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Commit(ctx context.Context, in *PreparedTransaction, opts ...grpc.CallOption) (Gateway_CommitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gateway_serviceDesc.Streams[0], "/protos.Gateway/Commit", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayCommitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_CommitClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type gatewayCommitClient struct {
	grpc.ClientStream
}

func (x *gatewayCommitClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) Evaluate(ctx context.Context, in *ProposedTransaction, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/protos.Gateway/Evaluate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	Prepare(context.Context, *ProposedTransaction) (*PreparedTransaction, error)
	Commit(*PreparedTransaction, Gateway_CommitServer) error
	Evaluate(context.Context, *ProposedTransaction) (*Result, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) Prepare(context.Context, *ProposedTransaction) (*PreparedTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (*UnimplementedGatewayServer) Commit(*PreparedTransaction, Gateway_CommitServer) error {
	return status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (*UnimplementedGatewayServer) Evaluate(context.Context, *ProposedTransaction) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposedTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Gateway/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Prepare(ctx, req.(*ProposedTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Commit_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PreparedTransaction)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).Commit(m, &gatewayCommitServer{stream})
}

type Gateway_CommitServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type gatewayCommitServer struct {
	grpc.ServerStream
}

func (x *gatewayCommitServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposedTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Gateway/Evaluate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Evaluate(ctx, req.(*ProposedTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Gateway_Prepare_Handler,
		},
		{
			MethodName: "Evaluate",
			Handler:    _Gateway_Evaluate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Commit",
			Handler:       _Gateway_Commit_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/gateway.proto",
}
