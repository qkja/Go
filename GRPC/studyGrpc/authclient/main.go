package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	userproto "studyGrpc/proto"
	"sync/atomic"
)

//func main() {
//	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
//	if err != nil {
//		panic(err)
//	}
//	client := userproto.NewAuthServiceClient(cc)
//	resp, err := client.Login(context.Background(), &userproto.LoginRequest{
//		Username: "admin",
//		Password: "admin",
//	})
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(resp.Token, resp.User.Id)
//}
//
//// 下面说一下池化技术
//
//type userClientPool struct {
//	clients []userproto.AuthServiceClient
//	index   int64
//}

func main() {
	cc := NewUserClientPool("localhost:8080", 10)
	resp, err := cc.Get().Login(context.Background(), &userproto.LoginRequest{
		Username: "admin",
		Password: "admin",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Token, resp.User.Id)
}

// 下面说一下池化技术

type userClientPool struct {
	clients []userproto.AuthServiceClient
	index   int64
}

func NewUserClientPool(addr string, size int) *userClientPool {

	var clients []userproto.AuthServiceClient
	for i := 0; i < size; i++ {
		cc, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		clients = append(clients, userproto.NewAuthServiceClient(cc))
	}
	return &userClientPool{clients: clients, index: 0}
}
func (p *userClientPool) Get() userproto.AuthServiceClient {
	index := atomic.AddInt64(&p.index, 1)
	return p.clients[int(index)%len(p.clients)]
}
