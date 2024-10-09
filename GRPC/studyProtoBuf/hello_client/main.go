package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	service "studyProtoBuf/hello_server/proto"
)

func main() {

	creds, _ := credentials.NewClientTLSFromFile("D:\\Code\\Go\\GRPC\\studyProtoBuf\\key\\test.pem", "*.qkj.com")

	//连接到server端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect:#{err}")
	}

	defer conn.Close()
	// 建立连接
	client := service.NewSayHelloClient(conn)
	//执行rpc调用(这个方法在服务器端来实现并返回结果)
	resp, _ := client.SayHello(context.Background(), &service.HelloRequest{RequestName: "qkj"})
	fmt.Println(resp.GetResponseMsg())
}

//# 1、生成私钥
//openssl genrsa -out server.key 2048
//# 2、生成证书 全部回车即可，可以不填4
// openssl req -new -x509 -key server.key -out server.crt -days 36500
//#3、生成 csr
//openssl req -new -key server.key -out server.csr

//#生成证书私钥test.key
//openssl genpkey -algorithm RSA -out test.key
//
//4 #通过私钥test.key生成证书请求文件test.csr(注意cfg和cnf)
//openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config ./openssl.cfg -extensions v3_req
//#test.csr是上面生成的证书请求文件。ca.crt/server.key是CA证书文件和key，用来对test.csr进行签名认证。这两个文件在第一部分生成。
//#生成SAN证书 pem
//openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cfg -extensions v3_req
