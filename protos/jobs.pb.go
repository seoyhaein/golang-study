// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/jobs.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JobsResponse_Status int32

const (
	JobsResponse_IN_PROGRESS JobsResponse_Status = 0
	JobsResponse_CANCELED    JobsResponse_Status = 1
	JobsResponse_FAILED      JobsResponse_Status = 2
	JobsResponse_COMPLETED   JobsResponse_Status = 3
	JobsResponse_TIME_OUT    JobsResponse_Status = 4
	JobsResponse_FINISHED    JobsResponse_Status = 5
)

var JobsResponse_Status_name = map[int32]string{
	0: "IN_PROGRESS",
	1: "CANCELED",
	2: "FAILED",
	3: "COMPLETED",
	4: "TIME_OUT",
	5: "FINISHED",
}

var JobsResponse_Status_value = map[string]int32{
	"IN_PROGRESS": 0,
	"CANCELED":    1,
	"FAILED":      2,
	"COMPLETED":   3,
	"TIME_OUT":    4,
	"FINISHED":    5,
}

func (x JobsResponse_Status) String() string {
	return proto.EnumName(JobsResponse_Status_name, int32(x))
}

func (JobsResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54c98b8c82adc31b, []int{1, 0}
}

type JobsRequest struct {
	JobReqId  int64                     `protobuf:"varint,1,opt,name=job_req_id,json=jobReqId,proto3" json:"job_req_id,omitempty"`
	InputData map[string]*_struct.Value `protobuf:"bytes,2,rep,name=input_data,json=inputData,proto3" json:"input_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 여기에는 일단 shell script 가 들어가서 string 이 되어 야 하지만, 일단 any 로 한번 넣어 보자
	InputMessage         *any.Any `protobuf:"bytes,3,opt,name=input_message,json=inputMessage,proto3" json:"input_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobsRequest) Reset()         { *m = JobsRequest{} }
func (m *JobsRequest) String() string { return proto.CompactTextString(m) }
func (*JobsRequest) ProtoMessage()    {}
func (*JobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c98b8c82adc31b, []int{0}
}

func (m *JobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobsRequest.Unmarshal(m, b)
}
func (m *JobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobsRequest.Marshal(b, m, deterministic)
}
func (m *JobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobsRequest.Merge(m, src)
}
func (m *JobsRequest) XXX_Size() int {
	return xxx_messageInfo_JobsRequest.Size(m)
}
func (m *JobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobsRequest proto.InternalMessageInfo

func (m *JobsRequest) GetJobReqId() int64 {
	if m != nil {
		return m.JobReqId
	}
	return 0
}

func (m *JobsRequest) GetInputData() map[string]*_struct.Value {
	if m != nil {
		return m.InputData
	}
	return nil
}

func (m *JobsRequest) GetInputMessage() *any.Any {
	if m != nil {
		return m.InputMessage
	}
	return nil
}

type JobsResponse struct {
	JobReqId             int64                     `protobuf:"varint,1,opt,name=job_req_id,json=jobReqId,proto3" json:"job_req_id,omitempty"`
	Status               JobsResponse_Status       `protobuf:"varint,2,opt,name=status,proto3,enum=protos.JobsResponse_Status" json:"status,omitempty"`
	OutputData           map[string]*_struct.Value `protobuf:"bytes,3,rep,name=output_data,json=outputData,proto3" json:"output_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OutputMessage        *any.Any                  `protobuf:"bytes,4,opt,name=output_message,json=outputMessage,proto3" json:"output_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *JobsResponse) Reset()         { *m = JobsResponse{} }
func (m *JobsResponse) String() string { return proto.CompactTextString(m) }
func (*JobsResponse) ProtoMessage()    {}
func (*JobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c98b8c82adc31b, []int{1}
}

func (m *JobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobsResponse.Unmarshal(m, b)
}
func (m *JobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobsResponse.Marshal(b, m, deterministic)
}
func (m *JobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobsResponse.Merge(m, src)
}
func (m *JobsResponse) XXX_Size() int {
	return xxx_messageInfo_JobsResponse.Size(m)
}
func (m *JobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobsResponse proto.InternalMessageInfo

func (m *JobsResponse) GetJobReqId() int64 {
	if m != nil {
		return m.JobReqId
	}
	return 0
}

func (m *JobsResponse) GetStatus() JobsResponse_Status {
	if m != nil {
		return m.Status
	}
	return JobsResponse_IN_PROGRESS
}

func (m *JobsResponse) GetOutputData() map[string]*_struct.Value {
	if m != nil {
		return m.OutputData
	}
	return nil
}

func (m *JobsResponse) GetOutputMessage() *any.Any {
	if m != nil {
		return m.OutputMessage
	}
	return nil
}

// 향후 삭제
// test code for any.Any
type AnyString struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnyString) Reset()         { *m = AnyString{} }
func (m *AnyString) String() string { return proto.CompactTextString(m) }
func (*AnyString) ProtoMessage()    {}
func (*AnyString) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c98b8c82adc31b, []int{2}
}

func (m *AnyString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnyString.Unmarshal(m, b)
}
func (m *AnyString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnyString.Marshal(b, m, deterministic)
}
func (m *AnyString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyString.Merge(m, src)
}
func (m *AnyString) XXX_Size() int {
	return xxx_messageInfo_AnyString.Size(m)
}
func (m *AnyString) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyString.DiscardUnknown(m)
}

var xxx_messageInfo_AnyString proto.InternalMessageInfo

func (m *AnyString) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.JobsResponse_Status", JobsResponse_Status_name, JobsResponse_Status_value)
	proto.RegisterType((*JobsRequest)(nil), "protos.JobsRequest")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "protos.JobsRequest.InputDataEntry")
	proto.RegisterType((*JobsResponse)(nil), "protos.JobsResponse")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "protos.JobsResponse.OutputDataEntry")
	proto.RegisterType((*AnyString)(nil), "protos.AnyString")
}

func init() { proto.RegisterFile("protos/jobs.proto", fileDescriptor_54c98b8c82adc31b) }

var fileDescriptor_54c98b8c82adc31b = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0xc6, 0xb7, 0x74, 0x17, 0xe1, 0x2d, 0xb0, 0x75, 0xdc, 0x18, 0xc4, 0x3d, 0x10, 0xa2, 0x09,
	0x07, 0x53, 0x0c, 0x7b, 0x51, 0x3c, 0x21, 0x74, 0xb5, 0x1b, 0xfe, 0x6c, 0xa6, 0xe0, 0xc1, 0x83,
	0x64, 0x4a, 0xc7, 0xa6, 0x08, 0x33, 0xd0, 0x99, 0x6e, 0xd2, 0xaf, 0xe0, 0x67, 0xf0, 0xbb, 0x6a,
	0xda, 0xa1, 0xab, 0xbb, 0x12, 0x4d, 0x8c, 0xa7, 0xbe, 0xef, 0xbc, 0xbf, 0x79, 0xfa, 0xcc, 0xd3,
	0x0e, 0x3c, 0xdc, 0x46, 0x5c, 0x72, 0xd1, 0x59, 0x71, 0x4f, 0x58, 0x59, 0x8d, 0x8a, 0x6a, 0xa9,
	0x71, 0x1e, 0x70, 0x1e, 0xac, 0x69, 0x27, 0x6b, 0xbd, 0xf8, 0x73, 0x47, 0xc8, 0x28, 0x5e, 0x4a,
	0x45, 0x35, 0x9e, 0xdc, 0x9f, 0x12, 0x96, 0xa8, 0x51, 0xeb, 0xbb, 0x06, 0xc6, 0x15, 0xf7, 0x04,
	0xa6, 0xbb, 0x98, 0x0a, 0x89, 0xce, 0x01, 0x56, 0xdc, 0x5b, 0x44, 0x74, 0xb7, 0x08, 0xfd, 0xba,
	0xd6, 0xd4, 0xda, 0x3a, 0x2e, 0xad, 0xb8, 0x87, 0xe9, 0xce, 0xf1, 0x51, 0x1f, 0x20, 0x64, 0xdb,
	0x58, 0x2e, 0x7c, 0x22, 0x49, 0xbd, 0xd0, 0xd4, 0xdb, 0x46, 0xb7, 0xa5, 0x94, 0x84, 0xf5, 0x8b,
	0x8c, 0xe5, 0xa4, 0xd4, 0x90, 0x48, 0x62, 0x33, 0x19, 0x25, 0xb8, 0x1c, 0xe6, 0x3d, 0x7a, 0x0d,
	0x55, 0x25, 0xb1, 0xa1, 0x42, 0x90, 0x80, 0xd6, 0xf5, 0xa6, 0xd6, 0x36, 0xba, 0x67, 0x96, 0xf2,
	0x68, 0xe5, 0x1e, 0xad, 0x3e, 0x4b, 0x70, 0x25, 0x43, 0xc7, 0x8a, 0x6c, 0xcc, 0xa0, 0x76, 0x57,
	0x17, 0x99, 0xa0, 0x7f, 0xa1, 0x49, 0x66, 0xb3, 0x8c, 0xd3, 0x12, 0xbd, 0x80, 0x93, 0x1b, 0xb2,
	0x8e, 0x69, 0xbd, 0x90, 0xc9, 0x3e, 0xfe, 0x4d, 0xf6, 0x43, 0x3a, 0xc5, 0x0a, 0xea, 0x15, 0x5e,
	0x69, 0xad, 0x6f, 0x3a, 0x54, 0x94, 0x75, 0xb1, 0xe5, 0x4c, 0xd0, 0xbf, 0x44, 0x70, 0x01, 0x45,
	0x21, 0x89, 0x8c, 0x45, 0xf6, 0x86, 0x5a, 0xf7, 0xe9, 0xdd, 0xe3, 0x2b, 0x0d, 0xcb, 0xcd, 0x10,
	0xbc, 0x47, 0x91, 0x0d, 0x06, 0x8f, 0xe5, 0x6d, 0x70, 0x7a, 0x16, 0xdc, 0xb3, 0x83, 0x3b, 0xa7,
	0x19, 0xf7, 0x33, 0x3a, 0xe0, 0xb7, 0x0b, 0xe8, 0x0d, 0xd4, 0xf6, 0x32, 0x79, 0x78, 0xc7, 0x7f,
	0x08, 0xaf, 0xaa, 0xd8, 0x3c, 0xbd, 0x39, 0x9c, 0xde, 0xd3, 0xfe, 0x2f, 0xf1, 0x7d, 0x82, 0xa2,
	0x3a, 0x2c, 0x3a, 0x05, 0xc3, 0x99, 0x2c, 0xae, 0xf1, 0xf4, 0x1d, 0xb6, 0x5d, 0xd7, 0x3c, 0x42,
	0x15, 0x28, 0x0d, 0xfa, 0x93, 0x81, 0x3d, 0xb2, 0x87, 0xa6, 0x86, 0x00, 0x8a, 0x97, 0x7d, 0x27,
	0xad, 0x0b, 0xa8, 0x0a, 0xe5, 0xc1, 0x74, 0x7c, 0x3d, 0xb2, 0x67, 0xf6, 0xd0, 0xd4, 0x53, 0x70,
	0xe6, 0x8c, 0xed, 0xc5, 0x74, 0x3e, 0x33, 0x8f, 0xd3, 0xee, 0xd2, 0x99, 0x38, 0xee, 0x7b, 0x7b,
	0x68, 0x9e, 0xb4, 0x9e, 0x43, 0xb9, 0xcf, 0x12, 0x57, 0x46, 0x21, 0x0b, 0x50, 0x1d, 0x1e, 0x2c,
	0xf9, 0x66, 0x43, 0x98, 0xbf, 0x37, 0x9d, 0xb7, 0xdd, 0xaf, 0x1a, 0x98, 0x23, 0xce, 0x82, 0x51,
	0x78, 0x43, 0xfd, 0x2b, 0xee, 0x0d, 0xc8, 0x7a, 0x8d, 0x7a, 0x50, 0x76, 0x63, 0x4f, 0x2c, 0xa3,
	0xd0, 0xa3, 0xe8, 0xd1, 0x81, 0xff, 0xb4, 0x71, 0x76, 0xe8, 0x1b, 0xb4, 0x8e, 0x5e, 0x6a, 0xa8,
	0x07, 0xc6, 0x9c, 0x89, 0x7f, 0xda, 0xfd, 0xb6, 0xf4, 0x71, 0x7f, 0x2f, 0x3d, 0xf5, 0xbc, 0xf8,
	0x11, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa5, 0xf8, 0x60, 0xbb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LongLivedJobCallClient is the client API for LongLivedJobCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LongLivedJobCallClient interface {
	Subscribe(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (LongLivedJobCall_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsResponse, error)
}

type longLivedJobCallClient struct {
	cc *grpc.ClientConn
}

func NewLongLivedJobCallClient(cc *grpc.ClientConn) LongLivedJobCallClient {
	return &longLivedJobCallClient{cc}
}

func (c *longLivedJobCallClient) Subscribe(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (LongLivedJobCall_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LongLivedJobCall_serviceDesc.Streams[0], "/protos.LongLivedJobCall/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &longLivedJobCallSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LongLivedJobCall_SubscribeClient interface {
	Recv() (*JobsResponse, error)
	grpc.ClientStream
}

type longLivedJobCallSubscribeClient struct {
	grpc.ClientStream
}

func (x *longLivedJobCallSubscribeClient) Recv() (*JobsResponse, error) {
	m := new(JobsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *longLivedJobCallClient) Unsubscribe(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsResponse, error) {
	out := new(JobsResponse)
	err := c.cc.Invoke(ctx, "/protos.LongLivedJobCall/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LongLivedJobCallServer is the server API for LongLivedJobCall service.
type LongLivedJobCallServer interface {
	Subscribe(*JobsRequest, LongLivedJobCall_SubscribeServer) error
	Unsubscribe(context.Context, *JobsRequest) (*JobsResponse, error)
}

// UnimplementedLongLivedJobCallServer can be embedded to have forward compatible implementations.
type UnimplementedLongLivedJobCallServer struct {
}

func (*UnimplementedLongLivedJobCallServer) Subscribe(req *JobsRequest, srv LongLivedJobCall_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedLongLivedJobCallServer) Unsubscribe(ctx context.Context, req *JobsRequest) (*JobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}

func RegisterLongLivedJobCallServer(s *grpc.Server, srv LongLivedJobCallServer) {
	s.RegisterService(&_LongLivedJobCall_serviceDesc, srv)
}

func _LongLivedJobCall_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JobsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LongLivedJobCallServer).Subscribe(m, &longLivedJobCallSubscribeServer{stream})
}

type LongLivedJobCall_SubscribeServer interface {
	Send(*JobsResponse) error
	grpc.ServerStream
}

type longLivedJobCallSubscribeServer struct {
	grpc.ServerStream
}

func (x *longLivedJobCallSubscribeServer) Send(m *JobsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LongLivedJobCall_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LongLivedJobCallServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LongLivedJobCall/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LongLivedJobCallServer).Unsubscribe(ctx, req.(*JobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LongLivedJobCall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.LongLivedJobCall",
	HandlerType: (*LongLivedJobCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _LongLivedJobCall_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _LongLivedJobCall_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/jobs.proto",
}