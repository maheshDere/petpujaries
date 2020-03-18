// Code generated by protoc-gen-go. DO NOT EDIT.
// source: downloader.proto

package downloader

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EmployeeFileDownloadRequest struct {
	AdminID              uint64   `protobuf:"varint,1,opt,name=adminID,proto3" json:"adminID,omitempty"`
	TotalEmployeeCount   uint64   `protobuf:"varint,2,opt,name=total_employee_count,json=totalEmployeeCount,proto3" json:"total_employee_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmployeeFileDownloadRequest) Reset()         { *m = EmployeeFileDownloadRequest{} }
func (m *EmployeeFileDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*EmployeeFileDownloadRequest) ProtoMessage()    {}
func (*EmployeeFileDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{0}
}

func (m *EmployeeFileDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeFileDownloadRequest.Unmarshal(m, b)
}
func (m *EmployeeFileDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeFileDownloadRequest.Marshal(b, m, deterministic)
}
func (m *EmployeeFileDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeFileDownloadRequest.Merge(m, src)
}
func (m *EmployeeFileDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_EmployeeFileDownloadRequest.Size(m)
}
func (m *EmployeeFileDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeFileDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeFileDownloadRequest proto.InternalMessageInfo

func (m *EmployeeFileDownloadRequest) GetAdminID() uint64 {
	if m != nil {
		return m.AdminID
	}
	return 0
}

func (m *EmployeeFileDownloadRequest) GetTotalEmployeeCount() uint64 {
	if m != nil {
		return m.TotalEmployeeCount
	}
	return 0
}

type EmployeeFileDownloadResponse struct {
	EmployeeDetails      []*EmployeeData `protobuf:"bytes,2,rep,name=employee_details,json=employeeDetails,proto3" json:"employee_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EmployeeFileDownloadResponse) Reset()         { *m = EmployeeFileDownloadResponse{} }
func (m *EmployeeFileDownloadResponse) String() string { return proto.CompactTextString(m) }
func (*EmployeeFileDownloadResponse) ProtoMessage()    {}
func (*EmployeeFileDownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{1}
}

func (m *EmployeeFileDownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeFileDownloadResponse.Unmarshal(m, b)
}
func (m *EmployeeFileDownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeFileDownloadResponse.Marshal(b, m, deterministic)
}
func (m *EmployeeFileDownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeFileDownloadResponse.Merge(m, src)
}
func (m *EmployeeFileDownloadResponse) XXX_Size() int {
	return xxx_messageInfo_EmployeeFileDownloadResponse.Size(m)
}
func (m *EmployeeFileDownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeFileDownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeFileDownloadResponse proto.InternalMessageInfo

func (m *EmployeeFileDownloadResponse) GetEmployeeDetails() []*EmployeeData {
	if m != nil {
		return m.EmployeeDetails
	}
	return nil
}

type EmployeeData struct {
	EmployeeData         []string `protobuf:"bytes,1,rep,name=employee_data,json=employeeData,proto3" json:"employee_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmployeeData) Reset()         { *m = EmployeeData{} }
func (m *EmployeeData) String() string { return proto.CompactTextString(m) }
func (*EmployeeData) ProtoMessage()    {}
func (*EmployeeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a99ec95c7ab1ff1, []int{2}
}

func (m *EmployeeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeData.Unmarshal(m, b)
}
func (m *EmployeeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeData.Marshal(b, m, deterministic)
}
func (m *EmployeeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeData.Merge(m, src)
}
func (m *EmployeeData) XXX_Size() int {
	return xxx_messageInfo_EmployeeData.Size(m)
}
func (m *EmployeeData) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeData.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeData proto.InternalMessageInfo

func (m *EmployeeData) GetEmployeeData() []string {
	if m != nil {
		return m.EmployeeData
	}
	return nil
}

func init() {
	proto.RegisterType((*EmployeeFileDownloadRequest)(nil), "downloader.EmployeeFileDownloadRequest")
	proto.RegisterType((*EmployeeFileDownloadResponse)(nil), "downloader.EmployeeFileDownloadResponse")
	proto.RegisterType((*EmployeeData)(nil), "downloader.EmployeeData")
}

func init() { proto.RegisterFile("downloader.proto", fileDescriptor_6a99ec95c7ab1ff1) }

var fileDescriptor_6a99ec95c7ab1ff1 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xdd, 0x56, 0x14, 0xc7, 0x4a, 0x4b, 0xf0, 0x10, 0xd4, 0xc3, 0xb2, 0x1e, 0xdc, 0x53,
	0x91, 0xf6, 0x27, 0x74, 0x15, 0xbc, 0xc6, 0x1f, 0x50, 0xc6, 0xcd, 0x1c, 0x02, 0x69, 0x66, 0xdd,
	0x4c, 0x15, 0xaf, 0xfe, 0x72, 0xe9, 0xb6, 0xb1, 0x2b, 0x2c, 0x78, 0xcc, 0xbc, 0xf7, 0xbe, 0xbc,
	0x49, 0x60, 0x66, 0xf9, 0x33, 0x78, 0x46, 0x4b, 0xed, 0xbc, 0x69, 0x59, 0x58, 0xc1, 0x71, 0x52,
	0x38, 0xb8, 0x7d, 0xda, 0x34, 0x9e, 0xbf, 0x88, 0x9e, 0x9d, 0xa7, 0xea, 0xa0, 0x18, 0x7a, 0xdf,
	0x52, 0x14, 0xa5, 0xe1, 0x1c, 0xed, 0xc6, 0x85, 0x97, 0x4a, 0x67, 0x79, 0x56, 0x9e, 0x9a, 0x74,
	0x54, 0x8f, 0x70, 0x2d, 0x2c, 0xe8, 0xd7, 0x74, 0x88, 0xaf, 0x6b, 0xde, 0x06, 0xd1, 0xa3, 0xce,
	0xa6, 0x3a, 0x2d, 0x91, 0x57, 0x3b, 0xa5, 0xa8, 0xe1, 0x6e, 0xf8, 0xaa, 0xd8, 0x70, 0x88, 0xa4,
	0x56, 0x30, 0xfb, 0x65, 0x59, 0x12, 0x74, 0x3e, 0xea, 0x51, 0x3e, 0x2e, 0x2f, 0x17, 0x7a, 0xde,
	0xdb, 0x21, 0x31, 0x2a, 0x14, 0x34, 0xd3, 0x94, 0xa8, 0xf6, 0x81, 0x62, 0x09, 0x93, 0xbe, 0x41,
	0xdd, 0xc3, 0xd5, 0x11, 0x8a, 0x82, 0x3a, 0xcb, 0xc7, 0xe5, 0x85, 0x99, 0x50, 0xcf, 0xb4, 0xf8,
	0xce, 0x60, 0x9a, 0xea, 0xbc, 0x52, 0xfb, 0xe1, 0x6a, 0x52, 0x0c, 0x3a, 0x8d, 0xfe, 0xb4, 0xde,
	0x41, 0x1f, 0x86, 0xfa, 0x0c, 0x3c, 0xdf, 0x4d, 0xf9, 0xbf, 0x71, 0xbf, 0x7c, 0x71, 0xf2, 0x76,
	0xd6, 0x7d, 0xce, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x83, 0xfb, 0x63, 0xb0, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownloadServiceClient is the client API for DownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownloadServiceClient interface {
	DownloadEmployeeFileData(ctx context.Context, in *EmployeeFileDownloadRequest, opts ...grpc.CallOption) (*EmployeeFileDownloadResponse, error)
}

type downloadServiceClient struct {
	cc *grpc.ClientConn
}

func NewDownloadServiceClient(cc *grpc.ClientConn) DownloadServiceClient {
	return &downloadServiceClient{cc}
}

func (c *downloadServiceClient) DownloadEmployeeFileData(ctx context.Context, in *EmployeeFileDownloadRequest, opts ...grpc.CallOption) (*EmployeeFileDownloadResponse, error) {
	out := new(EmployeeFileDownloadResponse)
	err := c.cc.Invoke(ctx, "/downloader.DownloadService/DownloadEmployeeFileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadServiceServer is the server API for DownloadService service.
type DownloadServiceServer interface {
	DownloadEmployeeFileData(context.Context, *EmployeeFileDownloadRequest) (*EmployeeFileDownloadResponse, error)
}

// UnimplementedDownloadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDownloadServiceServer struct {
}

func (*UnimplementedDownloadServiceServer) DownloadEmployeeFileData(ctx context.Context, req *EmployeeFileDownloadRequest) (*EmployeeFileDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadEmployeeFileData not implemented")
}

func RegisterDownloadServiceServer(s *grpc.Server, srv DownloadServiceServer) {
	s.RegisterService(&_DownloadService_serviceDesc, srv)
}

func _DownloadService_DownloadEmployeeFileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeFileDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).DownloadEmployeeFileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/downloader.DownloadService/DownloadEmployeeFileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).DownloadEmployeeFileData(ctx, req.(*EmployeeFileDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownloadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "downloader.DownloadService",
	HandlerType: (*DownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadEmployeeFileData",
			Handler:    _DownloadService_DownloadEmployeeFileData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "downloader.proto",
}