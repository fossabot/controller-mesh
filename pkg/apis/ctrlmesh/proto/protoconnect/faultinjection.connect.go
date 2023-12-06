// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pkg/apis/ctrlmesh/proto/faultinjection.proto

package protoconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	proto "github.com/KusionStack/controller-mesh/pkg/apis/ctrlmesh/proto"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// FaultInjectName is the fully-qualified name of the FaultInject service.
	FaultInjectName = "proto.FaultInject"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FaultInjectSendConfigProcedure is the fully-qualified name of the FaultInject's SendConfig RPC.
	FaultInjectSendConfigProcedure = "/proto.FaultInject/SendConfig"
)

// FaultInjectClient is a client for the proto.FaultInject service.
type FaultInjectClient interface {
	SendConfig(context.Context, *connect.Request[proto.FaultInjection]) (*connect.Response[proto.InjectResp], error)
}

// NewFaultInjectClient constructs a client for the proto.FaultInject service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFaultInjectClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FaultInjectClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &faultInjectClient{
		sendConfig: connect.NewClient[proto.FaultInjection, proto.InjectResp](
			httpClient,
			baseURL+FaultInjectSendConfigProcedure,
			opts...,
		),
	}
}

// faultInjectClient implements FaultInjectClient.
type faultInjectClient struct {
	sendConfig *connect.Client[proto.FaultInjection, proto.InjectResp]
}

// SendConfig calls proto.FaultInject.SendConfig.
func (c *faultInjectClient) SendConfig(ctx context.Context, req *connect.Request[proto.FaultInjection]) (*connect.Response[proto.InjectResp], error) {
	return c.sendConfig.CallUnary(ctx, req)
}

// FaultInjectHandler is an implementation of the proto.FaultInject service.
type FaultInjectHandler interface {
	SendConfig(context.Context, *connect.Request[proto.FaultInjection]) (*connect.Response[proto.InjectResp], error)
}

// NewFaultInjectHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFaultInjectHandler(svc FaultInjectHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	faultInjectSendConfigHandler := connect.NewUnaryHandler(
		FaultInjectSendConfigProcedure,
		svc.SendConfig,
		opts...,
	)
	return "/proto.FaultInject/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FaultInjectSendConfigProcedure:
			faultInjectSendConfigHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFaultInjectHandler returns CodeUnimplemented from all methods.
type UnimplementedFaultInjectHandler struct{}

func (UnimplementedFaultInjectHandler) SendConfig(context.Context, *connect.Request[proto.FaultInjection]) (*connect.Response[proto.InjectResp], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.FaultInject.SendConfig is not implemented"))
}
