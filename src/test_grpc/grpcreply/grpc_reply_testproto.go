//type GreeterServer interface {
//	// Sends a greeting
//	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
//	mustEmbedUnimplementedGreeterServer()
//}

package grpcreply

import (
	"context"
	"fmt"
	"test_grpc/proto/testproto"
)

type GrpcSayhelloServer struct {
	testproto.UnimplementedGreeterServer
}

func (*GrpcSayhelloServer) SayHello(ctx context.Context, req *testproto.HelloRequest) (*testproto.HelloReply, error) {

	fmt.Println("recv helloRequest name:", req.Name)
	rep := &testproto.HelloReply{}
	rep.Message = "sayhello reply ok"
	return rep, nil
}
