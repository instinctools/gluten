// Code generated by protoc-gen-go.
// source: rpc/master/rpcService.proto
// DO NOT EDIT!

/*
Package slave is a generated protocol buffer package.

It is generated from these files:
	rpc/master/rpcService.proto

It has these top-level messages:
	Step
	Execution
	ResponseMessage
*/
package slave

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

type Step struct {
	Name       string            `protobuf:"bytes,1,opt" json:"Name,omitempty"`
	Type       string            `protobuf:"bytes,2,opt" json:"Type,omitempty"`
	Parameters map[string]string `protobuf:"bytes,3,rep" json:"Parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SubSteps   []*Step           `protobuf:"bytes,4,rep" json:"SubSteps,omitempty"`
	Exec       *Execution        `protobuf:"bytes,5,opt" json:"Exec,omitempty"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}

func (m *Step) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Step) GetSubSteps() []*Step {
	if m != nil {
		return m.SubSteps
	}
	return nil
}

func (m *Step) GetExec() *Execution {
	if m != nil {
		return m.Exec
	}
	return nil
}

type Execution struct {
	ID     string `protobuf:"bytes,1,opt" json:"ID,omitempty"`
	Status string `protobuf:"bytes,2,opt" json:"Status,omitempty"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}

type ResponseMessage struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Step)(nil), "proto_service.Step")
	proto.RegisterType((*ResponseMessage)(nil), "proto_service.ResponseMessage")
	proto.RegisterType((*Execution)(nil), "proto_service.Execution")
}

// Client API for ProtoService service

type ProtoServiceClient interface {
	SendMessage(ctx context.Context, in *Step, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type protoServiceClient struct {
	cc *grpc.ClientConn
}

func NewProtoServiceClient(cc *grpc.ClientConn) ProtoServiceClient {
	return &protoServiceClient{cc}
}

func (c *protoServiceClient) SendMessage(ctx context.Context, in *Step, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := grpc.Invoke(ctx, "/slave.ProtoService/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtoService service

type ProtoServiceServer interface {
	SendMessage(context.Context, *Step) (*ResponseMessage, error)
}

func RegisterProtoServiceServer(s *grpc.Server, srv ProtoServiceServer) {
	s.RegisterService(&_ProtoService_serviceDesc, srv)
}

func _ProtoService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Step)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ProtoService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoServiceServer).SendMessage(ctx, req.(*Step))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slave.ProtoService",
	HandlerType: (*ProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ProtoService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}