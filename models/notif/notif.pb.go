// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notif.proto

package notif

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TopicType int32

const (
	TopicType_ALERT TopicType = 0
	TopicType_MODE  TopicType = 1
	TopicType_TEMP  TopicType = 2
)

var TopicType_name = map[int32]string{
	0: "ALERT",
	1: "MODE",
	2: "TEMP",
}

var TopicType_value = map[string]int32{
	"ALERT": 0,
	"MODE":  1,
	"TEMP":  2,
}

func (x TopicType) String() string {
	return proto.EnumName(TopicType_name, int32(x))
}

func (TopicType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{0}
}

type Alert struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alert) Reset()         { *m = Alert{} }
func (m *Alert) String() string { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()    {}
func (*Alert) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{0}
}

func (m *Alert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alert.Unmarshal(m, b)
}
func (m *Alert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alert.Marshal(b, m, deterministic)
}
func (m *Alert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alert.Merge(m, src)
}
func (m *Alert) XXX_Size() int {
	return xxx_messageInfo_Alert.Size(m)
}
func (m *Alert) XXX_DiscardUnknown() {
	xxx_messageInfo_Alert.DiscardUnknown(m)
}

var xxx_messageInfo_Alert proto.InternalMessageInfo

func (m *Alert) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Mode struct {
	NewMode              string   `protobuf:"bytes,1,opt,name=newMode,proto3" json:"newMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mode) Reset()         { *m = Mode{} }
func (m *Mode) String() string { return proto.CompactTextString(m) }
func (*Mode) ProtoMessage()    {}
func (*Mode) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{1}
}

func (m *Mode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mode.Unmarshal(m, b)
}
func (m *Mode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mode.Marshal(b, m, deterministic)
}
func (m *Mode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mode.Merge(m, src)
}
func (m *Mode) XXX_Size() int {
	return xxx_messageInfo_Mode.Size(m)
}
func (m *Mode) XXX_DiscardUnknown() {
	xxx_messageInfo_Mode.DiscardUnknown(m)
}

var xxx_messageInfo_Mode proto.InternalMessageInfo

func (m *Mode) GetNewMode() string {
	if m != nil {
		return m.NewMode
	}
	return ""
}

type Temp struct {
	NewTemp              uint32   `protobuf:"varint,1,opt,name=newTemp,proto3" json:"newTemp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Temp) Reset()         { *m = Temp{} }
func (m *Temp) String() string { return proto.CompactTextString(m) }
func (*Temp) ProtoMessage()    {}
func (*Temp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{2}
}

func (m *Temp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Temp.Unmarshal(m, b)
}
func (m *Temp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Temp.Marshal(b, m, deterministic)
}
func (m *Temp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Temp.Merge(m, src)
}
func (m *Temp) XXX_Size() int {
	return xxx_messageInfo_Temp.Size(m)
}
func (m *Temp) XXX_DiscardUnknown() {
	xxx_messageInfo_Temp.DiscardUnknown(m)
}

var xxx_messageInfo_Temp proto.InternalMessageInfo

func (m *Temp) GetNewTemp() uint32 {
	if m != nil {
		return m.NewTemp
	}
	return 0
}

type RegistrationRequest struct {
	ClientName           string   `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationRequest) Reset()         { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()    {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{3}
}

func (m *RegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationRequest.Unmarshal(m, b)
}
func (m *RegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationRequest.Marshal(b, m, deterministic)
}
func (m *RegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationRequest.Merge(m, src)
}
func (m *RegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrationRequest.Size(m)
}
func (m *RegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationRequest proto.InternalMessageInfo

func (m *RegistrationRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type RegistrationResponse struct {
	ClientName           string   `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	ServerName           string   `protobuf:"bytes,2,opt,name=serverName,proto3" json:"serverName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationResponse) Reset()         { *m = RegistrationResponse{} }
func (m *RegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*RegistrationResponse) ProtoMessage()    {}
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{4}
}

func (m *RegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationResponse.Unmarshal(m, b)
}
func (m *RegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationResponse.Marshal(b, m, deterministic)
}
func (m *RegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationResponse.Merge(m, src)
}
func (m *RegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_RegistrationResponse.Size(m)
}
func (m *RegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationResponse proto.InternalMessageInfo

func (m *RegistrationResponse) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *RegistrationResponse) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

type Topic struct {
	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	// uint64 clientCookie = 1;
	Type                 TopicType `protobuf:"varint,2,opt,name=type,proto3,enum=notif.TopicType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{5}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Topic) GetType() TopicType {
	if m != nil {
		return m.Type
	}
	return TopicType_ALERT
}

type Notification struct {
	ClientName           string    `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	ServerName           string    `protobuf:"bytes,2,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Type                 TopicType `protobuf:"varint,3,opt,name=type,proto3,enum=notif.TopicType" json:"type,omitempty"`
	A                    *Alert    `protobuf:"bytes,6,opt,name=a,proto3" json:"a,omitempty"`
	M                    *Mode     `protobuf:"bytes,4,opt,name=m,proto3" json:"m,omitempty"`
	T                    *Temp     `protobuf:"bytes,5,opt,name=t,proto3" json:"t,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{6}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Notification) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *Notification) GetType() TopicType {
	if m != nil {
		return m.Type
	}
	return TopicType_ALERT
}

func (m *Notification) GetA() *Alert {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Notification) GetM() *Mode {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Notification) GetT() *Temp {
	if m != nil {
		return m.T
	}
	return nil
}

func init() {
	proto.RegisterEnum("notif.TopicType", TopicType_name, TopicType_value)
	proto.RegisterType((*Alert)(nil), "notif.Alert")
	proto.RegisterType((*Mode)(nil), "notif.Mode")
	proto.RegisterType((*Temp)(nil), "notif.Temp")
	proto.RegisterType((*RegistrationRequest)(nil), "notif.RegistrationRequest")
	proto.RegisterType((*RegistrationResponse)(nil), "notif.RegistrationResponse")
	proto.RegisterType((*Topic)(nil), "notif.Topic")
	proto.RegisterType((*Notification)(nil), "notif.Notification")
}

func init() { proto.RegisterFile("notif.proto", fileDescriptor_bc3a72a17febd07f) }

var fileDescriptor_bc3a72a17febd07f = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x5f, 0x4f, 0xf2, 0x30,
	0x14, 0xc6, 0xdf, 0xf2, 0x6e, 0xbc, 0xec, 0xc0, 0x6b, 0x96, 0xe2, 0xc5, 0x9c, 0x89, 0xc1, 0xc5,
	0x0b, 0x42, 0x0c, 0x21, 0x33, 0x7e, 0x00, 0xa2, 0xdc, 0x39, 0x34, 0x75, 0xf1, 0x7e, 0xcc, 0x23,
	0x2e, 0x61, 0x7f, 0x5c, 0x8b, 0x86, 0xaf, 0xe5, 0x87, 0xf0, 0x73, 0x99, 0x76, 0x1d, 0x42, 0x42,
	0xc4, 0xc4, 0xbb, 0x3e, 0xfd, 0x3d, 0xe7, 0x9c, 0xb6, 0x4f, 0xa1, 0x9d, 0xe5, 0x22, 0x79, 0x1a,
	0x16, 0x65, 0x2e, 0x72, 0x6a, 0x2a, 0xe1, 0x9d, 0x82, 0x39, 0x5e, 0x60, 0x29, 0xa8, 0x03, 0xff,
	0x52, 0xe4, 0x3c, 0x9a, 0xa3, 0x43, 0x7a, 0xa4, 0x6f, 0xb1, 0x5a, 0x7a, 0x3d, 0x30, 0x82, 0xfc,
	0x11, 0xa5, 0x23, 0xc3, 0x37, 0xb9, 0xac, 0x1d, 0x5a, 0x4a, 0x47, 0x88, 0x69, 0xa1, 0x1d, 0x72,
	0xa9, 0x1c, 0xff, 0x59, 0x2d, 0xbd, 0x4b, 0xe8, 0x32, 0x9c, 0x27, 0x5c, 0x94, 0x91, 0x48, 0xf2,
	0x8c, 0xe1, 0xcb, 0x12, 0xb9, 0xa0, 0x27, 0x00, 0xf1, 0x22, 0xc1, 0x4c, 0x4c, 0xa3, 0xb4, 0xee,
	0xba, 0xb1, 0xe3, 0x3d, 0xc0, 0xe1, 0x76, 0x19, 0x2f, 0xf2, 0x8c, 0xe3, 0xbe, 0x3a, 0xc9, 0x39,
	0x96, 0xaf, 0x58, 0x2a, 0xde, 0xa8, 0xf8, 0xd7, 0x8e, 0x17, 0x80, 0x19, 0xe6, 0x45, 0x12, 0xef,
	0x6d, 0x74, 0x06, 0x86, 0x58, 0x15, 0x55, 0x8b, 0x03, 0xdf, 0x1e, 0x56, 0x2f, 0x28, 0x64, 0x6d,
	0xb8, 0x2a, 0x90, 0x29, 0xea, 0x7d, 0x10, 0xe8, 0x4c, 0x25, 0x49, 0x62, 0x75, 0xce, 0xdf, 0x9e,
	0x6f, 0x3d, 0xf6, 0xef, 0x77, 0x63, 0xa9, 0x0b, 0x24, 0x72, 0x9a, 0x3d, 0xd2, 0x6f, 0xfb, 0x1d,
	0x6d, 0x51, 0x59, 0x32, 0x12, 0xd1, 0x23, 0x20, 0xa9, 0x63, 0x28, 0xd6, 0xd6, 0x4c, 0x46, 0xc5,
	0x48, 0x2a, 0x91, 0x70, 0xcc, 0x2d, 0x24, 0x33, 0x62, 0x44, 0x0c, 0x06, 0x60, 0xad, 0x87, 0x50,
	0x0b, 0xcc, 0xf1, 0xcd, 0x84, 0x85, 0xf6, 0x1f, 0xda, 0x02, 0x23, 0xb8, 0xbd, 0x9e, 0xd8, 0x44,
	0xae, 0xc2, 0x49, 0x70, 0x67, 0x37, 0xfc, 0x77, 0x02, 0x56, 0xb1, 0xe4, 0xcf, 0xea, 0xe2, 0xf4,
	0x0a, 0x5a, 0x55, 0x52, 0x58, 0x52, 0x57, 0x77, 0xdd, 0x91, 0xb8, 0x7b, 0xbc, 0x93, 0xe9, 0x58,
	0xcf, 0xeb, 0xcf, 0x58, 0x5f, 0x47, 0x85, 0xe4, 0x76, 0xb5, 0xda, 0x7c, 0xe2, 0x11, 0xa1, 0x3e,
	0x58, 0xf7, 0xcb, 0x19, 0x8f, 0xcb, 0x64, 0x86, 0x3f, 0xa8, 0xe8, 0x93, 0x11, 0x99, 0x35, 0xd5,
	0xe7, 0xbf, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x18, 0xb2, 0x79, 0xd5, 0x0b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushNotifClient is the client API for PushNotif service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushNotifClient interface {
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	Alert(ctx context.Context, in *Topic, opts ...grpc.CallOption) (PushNotif_AlertClient, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (PushNotif_SubscribeClient, error)
}

type pushNotifClient struct {
	cc *grpc.ClientConn
}

func NewPushNotifClient(cc *grpc.ClientConn) PushNotifClient {
	return &pushNotifClient{cc}
}

func (c *pushNotifClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/notif.pushNotif/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushNotifClient) Alert(ctx context.Context, in *Topic, opts ...grpc.CallOption) (PushNotif_AlertClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushNotif_serviceDesc.Streams[0], "/notif.pushNotif/Alert", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushNotifAlertClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushNotif_AlertClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type pushNotifAlertClient struct {
	grpc.ClientStream
}

func (x *pushNotifAlertClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushNotifClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (PushNotif_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushNotif_serviceDesc.Streams[1], "/notif.pushNotif/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushNotifSubscribeClient{stream}
	return x, nil
}

type PushNotif_SubscribeClient interface {
	Send(*Topic) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type pushNotifSubscribeClient struct {
	grpc.ClientStream
}

func (x *pushNotifSubscribeClient) Send(m *Topic) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pushNotifSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushNotifServer is the server API for PushNotif service.
type PushNotifServer interface {
	Register(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	Alert(*Topic, PushNotif_AlertServer) error
	Subscribe(PushNotif_SubscribeServer) error
}

func RegisterPushNotifServer(s *grpc.Server, srv PushNotifServer) {
	s.RegisterService(&_PushNotif_serviceDesc, srv)
}

func _PushNotif_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushNotifServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notif.pushNotif/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushNotifServer).Register(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushNotif_Alert_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushNotifServer).Alert(m, &pushNotifAlertServer{stream})
}

type PushNotif_AlertServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type pushNotifAlertServer struct {
	grpc.ServerStream
}

func (x *pushNotifAlertServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _PushNotif_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PushNotifServer).Subscribe(&pushNotifSubscribeServer{stream})
}

type PushNotif_SubscribeServer interface {
	Send(*Notification) error
	Recv() (*Topic, error)
	grpc.ServerStream
}

type pushNotifSubscribeServer struct {
	grpc.ServerStream
}

func (x *pushNotifSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pushNotifSubscribeServer) Recv() (*Topic, error) {
	m := new(Topic)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PushNotif_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notif.pushNotif",
	HandlerType: (*PushNotifServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PushNotif_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Alert",
			Handler:       _PushNotif_Alert_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _PushNotif_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "notif.proto",
}
