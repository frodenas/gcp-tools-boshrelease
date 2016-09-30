// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/agent_service.proto
// DO NOT EDIT!

package google_monitoring_v3 // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_api4 "google.golang.org/genproto/googleapis/api/monitoredres"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The `CreateCollectdTimeSeries` request.
type CreateCollectdTimeSeriesRequest struct {
	// The project in which to create the time series. The format is
	// `"projects/PROJECT_ID_OR_NUMBER"`.
	Name string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	// The monitored resource associated with the time series.
	Resource *google_api4.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// The version of `collectd` that collected the data. Example: `"5.3.0-192.el6"`.
	CollectdVersion string `protobuf:"bytes,3,opt,name=collectd_version,json=collectdVersion" json:"collectd_version,omitempty"`
	// The `collectd` payloads representing the time series data.
	// You must not include more than a single point for each
	// time series, so no two payloads can have the same values
	// for all of the fields `plugin`, `plugin_instance`, `type`, and `type_instance`.
	CollectdPayloads []*CollectdPayload `protobuf:"bytes,4,rep,name=collectd_payloads,json=collectdPayloads" json:"collectd_payloads,omitempty"`
}

func (m *CreateCollectdTimeSeriesRequest) Reset()                    { *m = CreateCollectdTimeSeriesRequest{} }
func (m *CreateCollectdTimeSeriesRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCollectdTimeSeriesRequest) ProtoMessage()               {}
func (*CreateCollectdTimeSeriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateCollectdTimeSeriesRequest) GetResource() *google_api4.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *CreateCollectdTimeSeriesRequest) GetCollectdPayloads() []*CollectdPayload {
	if m != nil {
		return m.CollectdPayloads
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCollectdTimeSeriesRequest)(nil), "google.monitoring.v3.CreateCollectdTimeSeriesRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AgentTranslationService service

type AgentTranslationServiceClient interface {
	// **Stackdriver Monitoring Agent only:** Creates a new time series.
	//
	// <aside class="caution">This method is only for use by the Google Monitoring Agent.
	// Use [projects.timeSeries.create][google.monitoring.v3.MetricService.CreateTimeSeries]
	// instead.</aside>
	CreateCollectdTimeSeries(ctx context.Context, in *CreateCollectdTimeSeriesRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
}

type agentTranslationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentTranslationServiceClient(cc *grpc.ClientConn) AgentTranslationServiceClient {
	return &agentTranslationServiceClient{cc}
}

func (c *agentTranslationServiceClient) CreateCollectdTimeSeries(ctx context.Context, in *CreateCollectdTimeSeriesRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.monitoring.v3.AgentTranslationService/CreateCollectdTimeSeries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AgentTranslationService service

type AgentTranslationServiceServer interface {
	// **Stackdriver Monitoring Agent only:** Creates a new time series.
	//
	// <aside class="caution">This method is only for use by the Google Monitoring Agent.
	// Use [projects.timeSeries.create][google.monitoring.v3.MetricService.CreateTimeSeries]
	// instead.</aside>
	CreateCollectdTimeSeries(context.Context, *CreateCollectdTimeSeriesRequest) (*google_protobuf4.Empty, error)
}

func RegisterAgentTranslationServiceServer(s *grpc.Server, srv AgentTranslationServiceServer) {
	s.RegisterService(&_AgentTranslationService_serviceDesc, srv)
}

func _AgentTranslationService_CreateCollectdTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectdTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTranslationServiceServer).CreateCollectdTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.monitoring.v3.AgentTranslationService/CreateCollectdTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTranslationServiceServer).CreateCollectdTimeSeries(ctx, req.(*CreateCollectdTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentTranslationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.monitoring.v3.AgentTranslationService",
	HandlerType: (*AgentTranslationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollectdTimeSeries",
			Handler:    _AgentTranslationService_CreateCollectdTimeSeries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/agent_service.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x51, 0x41, 0x6b, 0x14, 0x31,
	0x18, 0x65, 0xda, 0x2a, 0x9a, 0x1e, 0xb4, 0x41, 0x74, 0x58, 0x10, 0x97, 0x05, 0x61, 0x2c, 0x98,
	0xc0, 0x0e, 0x1e, 0x14, 0x44, 0x6c, 0x11, 0xf1, 0x20, 0x2e, 0xd3, 0xe2, 0x75, 0xc9, 0x66, 0xbf,
	0xc6, 0xc8, 0x4c, 0xbe, 0x98, 0x64, 0x06, 0x16, 0xf1, 0xe2, 0x5f, 0xf0, 0xee, 0x0f, 0xf1, 0x6f,
	0xf8, 0x17, 0xbc, 0xfb, 0x17, 0x64, 0x26, 0x99, 0xed, 0xc1, 0x56, 0x8a, 0x97, 0x61, 0xf2, 0xe5,
	0xe5, 0xbd, 0xef, 0xbd, 0x47, 0x5e, 0x2b, 0x44, 0x55, 0x03, 0x53, 0x58, 0x0b, 0xa3, 0x18, 0x3a,
	0xc5, 0x15, 0x18, 0xeb, 0x30, 0x20, 0x8f, 0x57, 0xc2, 0x6a, 0xcf, 0x1b, 0x34, 0x3a, 0xa0, 0xd3,
	0x46, 0xf1, 0xae, 0xe4, 0x42, 0x81, 0x09, 0x4b, 0x0f, 0xae, 0xd3, 0x12, 0xd8, 0x00, 0xa6, 0x77,
	0x12, 0xd1, 0x39, 0x92, 0x75, 0xe5, 0xe4, 0xcd, 0xd5, 0xe8, 0x85, 0xd5, 0x3c, 0xd1, 0x49, 0x34,
	0x67, 0x5a, 0x71, 0x61, 0x0c, 0x06, 0x11, 0x34, 0x1a, 0x1f, 0x05, 0x26, 0xef, 0xae, 0x4e, 0x95,
	0x76, 0x80, 0xb5, 0x03, 0x7f, 0x7e, 0x58, 0x3a, 0xf0, 0xd8, 0xba, 0x71, 0xe3, 0xc9, 0x8b, 0xff,
	0xb6, 0x9e, 0x08, 0x4a, 0xa5, 0xc3, 0x87, 0x76, 0xc5, 0x24, 0x36, 0x3c, 0x92, 0xf0, 0xe1, 0x62,
	0xd5, 0x9e, 0x71, 0x1b, 0x36, 0x16, 0x3c, 0x87, 0xc6, 0x86, 0x4d, 0xfc, 0xc6, 0x47, 0xb3, 0xdf,
	0x19, 0x79, 0x70, 0xec, 0x40, 0x04, 0x38, 0xc6, 0xba, 0x06, 0x19, 0xd6, 0xa7, 0xba, 0x81, 0x13,
	0x70, 0x1a, 0x7c, 0x05, 0x9f, 0x5a, 0xf0, 0x81, 0x52, 0xb2, 0x67, 0x44, 0x03, 0xf9, 0xb5, 0x69,
	0x56, 0xdc, 0xac, 0x86, 0x7f, 0xfa, 0x94, 0xdc, 0x18, 0xf7, 0xcf, 0x77, 0xa6, 0x59, 0xb1, 0x3f,
	0xbf, 0xcf, 0x92, 0x01, 0x61, 0x35, 0x7b, 0x3b, 0xba, 0xac, 0x12, 0xa8, 0xda, 0xc2, 0xe9, 0x23,
	0x72, 0x5b, 0x26, 0xad, 0x65, 0x07, 0xce, 0x6b, 0x34, 0xf9, 0xee, 0x40, 0x7d, 0x6b, 0x9c, 0xbf,
	0x8f, 0x63, 0x5a, 0x91, 0x83, 0x2d, 0xd4, 0x8a, 0x4d, 0x8d, 0x62, 0xed, 0xf3, 0xbd, 0xe9, 0x6e,
	0xb1, 0x3f, 0x7f, 0xc8, 0x2e, 0x6a, 0x98, 0x8d, 0x2e, 0x16, 0x11, 0x5d, 0x6d, 0xa5, 0xd2, 0xc0,
	0xcf, 0x7f, 0x64, 0xe4, 0xde, 0xcb, 0x3e, 0xb6, 0x53, 0x27, 0x8c, 0xaf, 0x87, 0x52, 0x4f, 0x62,
	0xd9, 0xf4, 0x7b, 0x46, 0xf2, 0xcb, 0xd2, 0xa0, 0x4f, 0x2e, 0x51, 0xfc, 0x77, 0x7a, 0x93, 0xbb,
	0xe3, 0xb3, 0xb1, 0x0c, 0xf6, 0xaa, 0xcf, 0x7f, 0x56, 0x7e, 0xfd, 0xf9, 0xeb, 0xdb, 0xce, 0xe3,
	0x59, 0xd1, 0xf7, 0xf8, 0xb9, 0x0f, 0xf5, 0xb9, 0x75, 0xf8, 0x11, 0x64, 0xf0, 0xfc, 0xf0, 0x0b,
	0x97, 0x7f, 0x11, 0x3e, 0xcb, 0x0e, 0x8f, 0x38, 0xc9, 0x25, 0x36, 0x17, 0x2e, 0x72, 0x74, 0x30,
	0xb8, 0x4a, 0x56, 0x16, 0xbd, 0xd8, 0x22, 0x5b, 0x5d, 0x1f, 0x54, 0xcb, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x07, 0x6b, 0xcb, 0x33, 0x59, 0x03, 0x00, 0x00,
}