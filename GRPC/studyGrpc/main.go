package main

import "studyGrpc/internal/grpc"

func main() {
	// 1. 其启动服务
	grpc.StartGRPCServer()
}

// 测试HTTPde
//func main() {
//	var user = &userPorto.User{
//		Id:    1,
//		Name:  "abc",
//		Email: "qq.com",
//	}
//	// json 字符串
//	data, _ := json.Marshal(user)
//	fmt.Println(string(data))
//	// 自带的 序列化
//	msg, _ := proto.Marshal(user)
//	fmt.Println(string(msg))
//}
