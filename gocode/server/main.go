package main

import (
	"net"

	pb "testgrpc_com/xuexitest" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type testService struct{}

// HelloService ...
var TestService = testService{}

func (t testService) SayTest(ctx context.Context, in *pb.TestRequest) (*pb.TestReply, error) {
	resp := new(pb.TestReply)

	getdata1 := pb.TestReply_GetData{Id: 1, Name: "laixhe1"}
	getdata2 := pb.TestReply_GetData{Id: 2, Name: "laixhe2"}
	getdata3 := pb.TestReply_GetData{Id: 3, Name: "laixhe3"}
	resp.Getdataarr = append(resp.Getdataarr, &getdata1, &getdata2, &getdata3)

	grpclog.Println(in.Typeid)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterXuexitestServer(s, TestService)

	grpclog.Println("Listen on " + Address)

	s.Serve(listen)
}
