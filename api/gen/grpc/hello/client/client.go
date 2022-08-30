// Code generated by goa v3.8.4, DO NOT EDIT.
//
// hello gRPC client
//
// Command:
// $ goa gen github.com/piotrostr/full-auto-gke/api/design

package client

import (
	"context"

	hellopb "github.com/piotrostr/full-auto-gke/api/gen/grpc/hello/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli hellopb.HelloClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the hello service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: hellopb.NewHelloClient(cc),
		opts:    opts,
	}
}

// SayHello calls the "SayHello" function in hellopb.HelloClient interface.
func (c *Client) SayHello() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSayHelloFunc(c.grpccli, c.opts...),
			EncodeSayHelloRequest,
			DecodeSayHelloResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
