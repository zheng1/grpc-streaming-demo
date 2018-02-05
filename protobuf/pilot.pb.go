// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pilot.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	pilot.proto

It has these top-level messages:
	Request
	Response
	RegisterProxy
	ReportCmdStatus
	RegisterMetadata
	StartConfd
	RegisterCmd
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Register *RegisterProxy   `protobuf:"bytes,1,opt,name=register" json:"register,omitempty"`
	Report   *ReportCmdStatus `protobuf:"bytes,2,opt,name=report" json:"report,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetRegister() *RegisterProxy {
	if m != nil {
		return m.Register
	}
	return nil
}

func (m *Request) GetReport() *ReportCmdStatus {
	if m != nil {
		return m.Report
	}
	return nil
}

type Response struct {
	Register *RegisterMetadata `protobuf:"bytes,1,opt,name=register" json:"register,omitempty"`
	Start    *StartConfd       `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	Cmd      *RegisterCmd      `protobuf:"bytes,3,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetRegister() *RegisterMetadata {
	if m != nil {
		return m.Register
	}
	return nil
}

func (m *Response) GetStart() *StartConfd {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Response) GetCmd() *RegisterCmd {
	if m != nil {
		return m.Cmd
	}
	return nil
}

type RegisterProxy struct {
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
}

func (m *RegisterProxy) Reset()                    { *m = RegisterProxy{} }
func (m *RegisterProxy) String() string            { return proto.CompactTextString(m) }
func (*RegisterProxy) ProtoMessage()               {}
func (*RegisterProxy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterProxy) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

type ReportCmdStatus struct {
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ReportCmdStatus) Reset()                    { *m = ReportCmdStatus{} }
func (m *ReportCmdStatus) String() string            { return proto.CompactTextString(m) }
func (*ReportCmdStatus) ProtoMessage()               {}
func (*ReportCmdStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReportCmdStatus) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ReportCmdStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RegisterMetadata struct {
	Metadata string `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *RegisterMetadata) Reset()                    { *m = RegisterMetadata{} }
func (m *RegisterMetadata) String() string            { return proto.CompactTextString(m) }
func (*RegisterMetadata) ProtoMessage()               {}
func (*RegisterMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterMetadata) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type StartConfd struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *StartConfd) Reset()                    { *m = StartConfd{} }
func (m *StartConfd) String() string            { return proto.CompactTextString(m) }
func (*StartConfd) ProtoMessage()               {}
func (*StartConfd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StartConfd) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type RegisterCmd struct {
	Ip  string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Cmd string `protobuf:"bytes,2,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *RegisterCmd) Reset()                    { *m = RegisterCmd{} }
func (m *RegisterCmd) String() string            { return proto.CompactTextString(m) }
func (*RegisterCmd) ProtoMessage()               {}
func (*RegisterCmd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegisterCmd) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RegisterCmd) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pilot.Request")
	proto.RegisterType((*Response)(nil), "pilot.Response")
	proto.RegisterType((*RegisterProxy)(nil), "pilot.RegisterProxy")
	proto.RegisterType((*ReportCmdStatus)(nil), "pilot.ReportCmdStatus")
	proto.RegisterType((*RegisterMetadata)(nil), "pilot.RegisterMetadata")
	proto.RegisterType((*StartConfd)(nil), "pilot.StartConfd")
	proto.RegisterType((*RegisterCmd)(nil), "pilot.RegisterCmd")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pilot service

type PilotClient interface {
	Flight(ctx context.Context, opts ...grpc.CallOption) (Pilot_FlightClient, error)
}

type pilotClient struct {
	cc *grpc.ClientConn
}

func NewPilotClient(cc *grpc.ClientConn) PilotClient {
	return &pilotClient{cc}
}

func (c *pilotClient) Flight(ctx context.Context, opts ...grpc.CallOption) (Pilot_FlightClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Pilot_serviceDesc.Streams[0], c.cc, "/pilot.Pilot/Flight", opts...)
	if err != nil {
		return nil, err
	}
	x := &pilotFlightClient{stream}
	return x, nil
}

type Pilot_FlightClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type pilotFlightClient struct {
	grpc.ClientStream
}

func (x *pilotFlightClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pilotFlightClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Pilot service

type PilotServer interface {
	Flight(Pilot_FlightServer) error
}

func RegisterPilotServer(s *grpc.Server, srv PilotServer) {
	s.RegisterService(&_Pilot_serviceDesc, srv)
}

func _Pilot_Flight_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PilotServer).Flight(&pilotFlightServer{stream})
}

type Pilot_FlightServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type pilotFlightServer struct {
	grpc.ServerStream
}

func (x *pilotFlightServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pilotFlightServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Pilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pilot.Pilot",
	HandlerType: (*PilotServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Flight",
			Handler:       _Pilot_Flight_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pilot.proto",
}

func init() { proto.RegisterFile("pilot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x6b, 0x32, 0x31,
	0x10, 0x86, 0xbf, 0x55, 0x5c, 0xd7, 0x59, 0x3e, 0xb5, 0x43, 0x51, 0x91, 0x42, 0x4b, 0x28, 0xd4,
	0xd3, 0x2a, 0x7a, 0xe9, 0x59, 0xa1, 0xe0, 0xa1, 0x20, 0xf1, 0xd6, 0x4b, 0x89, 0x26, 0xda, 0xa0,
	0xbb, 0xd9, 0x6e, 0x22, 0xb4, 0x7f, 0xa2, 0xbf, 0xb9, 0x6c, 0x36, 0x5d, 0xed, 0xf6, 0x36, 0xef,
	0xcc, 0x93, 0xcc, 0xbc, 0x93, 0x40, 0x98, 0xca, 0xa3, 0x32, 0x51, 0x9a, 0x29, 0xa3, 0xb0, 0x61,
	0x05, 0x39, 0x40, 0x93, 0x8a, 0xf7, 0x93, 0xd0, 0x06, 0x27, 0x10, 0x64, 0x62, 0x2f, 0xb5, 0x11,
	0xd9, 0xc0, 0xbb, 0xf3, 0x46, 0xe1, 0xf4, 0x3a, 0x2a, 0x4e, 0x50, 0x97, 0x5e, 0x65, 0xea, 0xe3,
	0x93, 0x96, 0x14, 0x46, 0xe0, 0x67, 0x22, 0x55, 0x99, 0x19, 0xd4, 0x2c, 0xdf, 0x2b, 0xf9, 0x3c,
	0xb9, 0x88, 0xf9, 0xda, 0x30, 0x73, 0xd2, 0xd4, 0x51, 0xe4, 0xcb, 0x83, 0x80, 0x0a, 0x9d, 0xaa,
	0x44, 0x0b, 0x9c, 0xfd, 0x69, 0xd7, 0xaf, 0xb4, 0x7b, 0x16, 0x86, 0x71, 0x66, 0xd8, 0x45, 0xc7,
	0x07, 0x68, 0x68, 0xc3, 0xca, 0x86, 0x57, 0xee, 0xc4, 0x3a, 0xcf, 0x2d, 0x54, 0xb2, 0xe3, 0xb4,
	0xa8, 0xe3, 0x3d, 0xd4, 0xb7, 0x31, 0x1f, 0xd4, 0x2d, 0x86, 0x95, 0x8b, 0x17, 0x31, 0xa7, 0x79,
	0x99, 0x4c, 0xe0, 0xff, 0x2f, 0x6f, 0x78, 0x0b, 0xa1, 0x4c, 0xb4, 0x61, 0xc9, 0x56, 0xbc, 0x4a,
	0x6e, 0xe7, 0x6a, 0x51, 0xf8, 0x49, 0x2d, 0x39, 0x99, 0x43, 0xa7, 0xe2, 0x0e, 0xfb, 0xd0, 0x34,
	0x4c, 0x1f, 0xce, 0xbc, 0x9f, 0xcb, 0x25, 0xc7, 0x1e, 0xf8, 0xda, 0x22, 0x76, 0xda, 0x16, 0x75,
	0x8a, 0x44, 0xd0, 0xad, 0x5a, 0xc4, 0x21, 0x04, 0xb1, 0x8b, 0xdd, 0x2d, 0xa5, 0x26, 0x37, 0x00,
	0x67, 0x83, 0xd8, 0x86, 0x9a, 0x4c, 0x1d, 0x53, 0x93, 0x29, 0x19, 0x43, 0x78, 0xe1, 0xab, 0x5a,
	0xc6, 0x6e, 0xb1, 0x88, 0x62, 0x82, 0x3c, 0x9c, 0x3e, 0x42, 0x63, 0x95, 0xaf, 0x03, 0xc7, 0xe0,
	0x3f, 0x1d, 0xe5, 0xfe, 0xcd, 0x60, 0xbb, 0x5c, 0x90, 0xfd, 0x0a, 0xc3, 0x4e, 0xa9, 0x8b, 0xc7,
	0x22, 0xff, 0x46, 0xde, 0xc4, 0x9b, 0xc3, 0x4b, 0x60, 0x3f, 0xcf, 0xe6, 0xb4, 0xdb, 0xf8, 0x36,
	0x9a, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x85, 0xc0, 0xaf, 0x3a, 0x55, 0x02, 0x00, 0x00,
}
