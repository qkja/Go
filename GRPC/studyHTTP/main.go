package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

func test1() {
	listen, _ := net.Listen("tcp", ":8080")
	conn, _ := listen.Accept()
	reader := bufio.NewReader(conn)
	request, err := http.ReadRequest(reader)
	fmt.Println(err)
	fmt.Println(request.URL, request.Method, request.Header)
	conn.Write([]byte("HTTP/1.1 200 0K\r\nContent-Type: text/plain\r\nContent-Length: 11\r\n\r\nhello world"))
}

func test2() {
	listen, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := listen.Accept()
		go func() {
			for {
				reader := bufio.NewReader(conn)
				request, err := http.ReadRequest(reader)
				fmt.Println(err)
				fmt.Println(request.URL, request.Method, request.Header)
				conn.Write([]byte("HTTP/1.1 280 0K\r\nContent-Type: text/plain\r\nContent-Length: 11\r\n\r\nhello world"))
			}
		}()
	}

}

func test03() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//fmt.Println(request.Method)
		//fmt.Println(request.URL)
		fmt.Printf("%v", request)
		writer.Write([]byte("hello world"))
	})
	http.HandleFunc("/aa", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		writer.Write([]byte("hello world2"))
	})
	http.ListenAndServe(":8080", nil)
}
func main() {
	//test1()
	//test2()
	test03()
}
