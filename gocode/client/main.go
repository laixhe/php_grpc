package main

import (
	pb "testgrpc_com/xuexitest" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "192.168.0.60:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewXuexitestClient(conn)

	// 调用方法
	reqBody := new(pb.TestRequest)
	reqBody.Typeid = 1
	r, err := c.SayTest(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(r.Getdataarr)
}
