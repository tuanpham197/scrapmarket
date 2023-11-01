package rpc

import (
	"context"
	pb "github.com/tuanpham197/test_repo"
)

type rpcClient struct {
	client pb.GreeterClient
}

func NewRPCClient(client pb.GreeterClient) *rpcClient {
	return &rpcClient{client: client}
}

func (r *rpcClient) SayHello(ctx context.Context, name string) (*pb.HelloReply, error) {
	result, err := r.client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return result, nil
}
