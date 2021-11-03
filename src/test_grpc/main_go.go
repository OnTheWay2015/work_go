package main

import (
	"context"
	"fmt"
	"net"
	"test_grpc/grpcreply"
	"test_grpc/proto/testproto"
	"time"

	"google.golang.org/grpc"
)

func grpc_client() {

	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println("conn err:", err.Error())
		return
	}

	client := testproto.NewGreeterClient(conn)
	msg := testproto.HelloRequest{}
	msg.Name = "sayhello request"
	reply, err := client.SayHello(context.Background(), &msg)
	if err != nil {
		fmt.Println("client.SayHello err:", err.Error())
		return
	}
	fmt.Println("client.SayHello reply ok reply message:", reply.Message)

}
func grpc_server() {
	srv := grpc.NewServer()
	ig := grpcreply.GrpcSayhelloServer{}
	testproto.RegisterGreeterServer(srv, &ig)
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("listen err:", err.Error())
		return
	}

	err = srv.Serve(listener)
	if err != nil {
		fmt.Println("failed to serve: ", err.Error())
		return
	}
}

func main() {
	fmt.Print("test_grpc main act")
	go grpc_server()
	time.Sleep(time.Second * 15)
	grpc_client()
	fmt.Print("test_grpc main act end")
}
