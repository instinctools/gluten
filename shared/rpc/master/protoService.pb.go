package master

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type RequestMessage struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *RequestMessage) Reset()         { *m = RequestMessage{} }
func (m *RequestMessage) String() string { return proto.CompactTextString(m) }
func (*RequestMessage) ProtoMessage()    {}

type ResponseMessage struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}

func init() {
	proto.RegisterType((*RequestMessage)(nil), "proto_service.RequestMessage")
	proto.RegisterType((*ResponseMessage)(nil), "proto_service.ResponseMessage")
}

// Client API for ProtoService service

type ProtoServiceClient interface {
	SendMessage(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type protoServiceClient struct {
	cc *grpc.ClientConn
}

func NewProtoServiceClient(cc *grpc.ClientConn) ProtoServiceClient {
	return &protoServiceClient{cc}
}

func (c *protoServiceClient) SendMessage(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := grpc.Invoke(ctx, "/master.ProtoService/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtoService service

type ProtoServiceServer interface {
	SendMessage(context.Context, *RequestMessage) (*ResponseMessage, error)
}

func RegisterProtoServiceServer(s *grpc.Server, srv ProtoServiceServer) {
	s.RegisterService(&_ProtoService_serviceDesc, srv)
}

func _ProtoService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.ProtoService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoServiceServer).SendMessage(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "master.ProtoService",
	HandlerType: (*ProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ProtoService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
