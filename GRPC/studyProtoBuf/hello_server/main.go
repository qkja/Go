package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"maps"
	"net"
	service "studyProtoBuf/hello_server/proto"
)

type server struct {
	service.UnimplementedSayHelloServer
}

func (server) SayHello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		ResponseMsg: "hello" + req.RequestName,
	}, nil
}

func main() {
	// 0. 开启密钥
	creds, _ := credentials.NewServerTLSFromFile("D:\\Code\\Go\\GRPC\\studyProtoBuf\\key\\test.pem", "D:\\Code\\Go\\GRPC\\studyProtoBuf\\key\\test.key")

	// 1. 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("创建端口失败 %v", err)
		return
	}
	// 2. 创建服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 3.将我们的服务注册到grpc
	service.RegisterSayHelloServer(grpcServer, &server{})

	// 4. 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("监听失败")
	}
	type Programmer struct {
		Name string
		City string
	}

	m1 := map[string]Programmer{
		"programmer-01": {Name: "陈明勇", City: "深圳"},
		"programmer-02": {Name: "张三", City: "广州"},
	}
	var m2 = maps.Clone(m1)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m2: %v\n", m2)
}
