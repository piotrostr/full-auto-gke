// Code generated by goa v3.8.4, DO NOT EDIT.
//
// hello gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/piotrostr/full-auto-gke/api/design

package client

import (
	"context"

	hellopb "github.com/piotrostr/full-auto-gke/api/gen/grpc/hello/pb"
	hello "github.com/piotrostr/full-auto-gke/api/gen/hello"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildSayHelloFunc builds the remote method to invoke for "hello" service
// "say-hello" endpoint.
func BuildSayHelloFunc(grpccli hellopb.HelloClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.SayHello(ctx, reqpb.(*hellopb.SayHelloRequest), opts...)
		}
		return grpccli.SayHello(ctx, &hellopb.SayHelloRequest{}, opts...)
	}
}

// EncodeSayHelloRequest encodes requests sent to hello say-hello endpoint.
func EncodeSayHelloRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*hello.SayHelloPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("hello", "say-hello", "*hello.SayHelloPayload", v)
	}
	return NewProtoSayHelloRequest(payload), nil
}

// DecodeSayHelloResponse decodes responses from the hello say-hello endpoint.
func DecodeSayHelloResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*hellopb.SayHelloResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("hello", "say-hello", "*hellopb.SayHelloResponse", v)
	}
	res := NewSayHelloResult(message)
	return res, nil
}
