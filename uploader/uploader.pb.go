// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uploader.proto

package uploader

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UploadFileRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadFileRequest_Info
	//	*UploadFileRequest_Chuckdata
	Data                 isUploadFileRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b055a52f625709c9, []int{0}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadFileRequest_Chuckdata struct {
	Chuckdata []byte `protobuf:"bytes,2,opt,name=chuckdata,proto3,oneof"`
}

func (*UploadFileRequest_Info) isUploadFileRequest_Data() {}

func (*UploadFileRequest_Chuckdata) isUploadFileRequest_Data() {}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadFileRequest) GetInfo() *FileInfo {
	if x, ok := m.GetData().(*UploadFileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadFileRequest) GetChuckdata() []byte {
	if x, ok := m.GetData().(*UploadFileRequest_Chuckdata); ok {
		return x.Chuckdata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadFileRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadFileRequest_Info)(nil),
		(*UploadFileRequest_Chuckdata)(nil),
	}
}

type FileInfo struct {
	Modulename           string   `protobuf:"bytes,1,opt,name=modulename,proto3" json:"modulename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b055a52f625709c9, []int{1}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetModulename() string {
	if m != nil {
		return m.Modulename
	}
	return ""
}

type UploadFileResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Size                 uint32   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b055a52f625709c9, []int{2}
}

func (m *UploadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponse.Unmarshal(m, b)
}
func (m *UploadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponse.Merge(m, src)
}
func (m *UploadFileResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponse.Size(m)
}
func (m *UploadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponse proto.InternalMessageInfo

func (m *UploadFileResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadFileResponse) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UploadFileResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*UploadFileRequest)(nil), "uploader.UploadFileRequest")
	proto.RegisterType((*FileInfo)(nil), "uploader.FileInfo")
	proto.RegisterType((*UploadFileResponse)(nil), "uploader.UploadFileResponse")
}

func init() { proto.RegisterFile("uploader.proto", fileDescriptor_b055a52f625709c9) }

var fileDescriptor_b055a52f625709c9 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x4d, 0x20, 0x0a, 0xed, 0x41, 0x91, 0xb8, 0x01, 0x45, 0x80, 0xaa, 0x2a, 0x93, 0xc5, 0xd0,
	0xa1, 0xfc, 0x01, 0x03, 0x2a, 0x62, 0x33, 0x62, 0x41, 0x2c, 0x26, 0xb9, 0x82, 0x45, 0x62, 0x87,
	0x9c, 0xcd, 0xc0, 0xd7, 0xa3, 0xba, 0x98, 0x44, 0xa2, 0x93, 0xef, 0x3d, 0xdf, 0xf3, 0x7b, 0x7e,
	0x70, 0xea, 0xbb, 0xc6, 0xaa, 0x9a, 0xfa, 0x65, 0xd7, 0x5b, 0x67, 0x71, 0x12, 0x71, 0x49, 0x70,
	0xf6, 0x14, 0xe6, 0x3b, 0xdd, 0x90, 0xa4, 0x4f, 0x4f, 0xec, 0x50, 0x40, 0xa6, 0xcd, 0xc6, 0x16,
	0xe9, 0x22, 0x15, 0xc7, 0x2b, 0x5c, 0xfe, 0xa9, 0xb7, 0x4b, 0xf7, 0x66, 0x63, 0xd7, 0x89, 0x0c,
	0x1b, 0x38, 0x87, 0x69, 0xf5, 0xee, 0xab, 0x8f, 0x5a, 0x39, 0x55, 0x1c, 0x2c, 0x52, 0x71, 0xb2,
	0x4e, 0xe4, 0x40, 0xdd, 0xe6, 0x90, 0x6d, 0xcf, 0xf2, 0x1a, 0x26, 0x51, 0x8b, 0x73, 0x80, 0xd6,
	0xd6, 0xbe, 0x21, 0xa3, 0x5a, 0x0a, 0x1e, 0x53, 0x39, 0x62, 0xca, 0x67, 0xc0, 0x71, 0x24, 0xee,
	0xac, 0x61, 0xc2, 0x02, 0x8e, 0x5a, 0x62, 0x56, 0x6f, 0x51, 0x12, 0x21, 0x9e, 0x43, 0xce, 0x4e,
	0x39, 0xcf, 0x21, 0xc0, 0x4c, 0xfe, 0x22, 0x44, 0xc8, 0x58, 0x7f, 0x53, 0x71, 0x18, 0xd8, 0x30,
	0xaf, 0x5e, 0x60, 0xb6, 0x7b, 0xfb, 0x91, 0xfa, 0x2f, 0x5d, 0x11, 0x3e, 0x00, 0x0c, 0x66, 0x78,
	0x39, 0x7c, 0xf5, 0x5f, 0x2b, 0x17, 0x57, 0xfb, 0x2f, 0x77, 0xf9, 0xca, 0x44, 0xa4, 0xaf, 0x79,
	0x68, 0xf7, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x26, 0xfb, 0x6f, 0x82, 0x6f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploadServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadFileClient, error)
}

type uploadServiceClient struct {
	cc *grpc.ClientConn
}

func NewUploadServiceClient(cc *grpc.ClientConn) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UploadService_serviceDesc.Streams[0], "/uploader.UploadService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadFileClient{stream}
	return x, nil
}

type UploadService_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
type UploadServiceServer interface {
	UploadFile(UploadService_UploadFileServer) error
}

// UnimplementedUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (*UnimplementedUploadServiceServer) UploadFile(srv UploadService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

func RegisterUploadServiceServer(s *grpc.Server, srv UploadServiceServer) {
	s.RegisterService(&_UploadService_serviceDesc, srv)
}

func _UploadService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).UploadFile(&uploadServiceUploadFileServer{stream})
}

type UploadService_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uploader.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _UploadService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "uploader.proto",
}
