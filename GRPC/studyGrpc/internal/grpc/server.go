package grpc

import (
	"google.golang.org/grpc"
	"net"
	"studyGrpc/internal/ctrl"
	userPorto "studyGrpc/proto"
)

func StartGRPCServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// new出一个grpc
	s := grpc.NewServer()
	//注册我们的服务
	authController := ctrl.NewAuthController()
	userPorto.RegisterAuthServiceServer(s, authController)
	// 启动服务
	err = s.Serve(ln)
	if err != nil {
		return
	}
}
