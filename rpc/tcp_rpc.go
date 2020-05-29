package main

import (
// 	"errors"
	"fmt"
        "strconv"
	"net"
	"net/rpc"
)

type HttpRpcServer struct {
	a int
}
func (s *HttpRpcServer) Print( a *int, o *string) error {
	fmt.Println("HttpRpcServer.Print:",*a)
	*o = strconv.Itoa(*a)
        return nil
}
func main() {
	fmt.Println("Main start...")
	server := new(HttpRpcServer)
	rpc.Register(server)
        tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234") // 这里指定协议为tcp
	if err != nil {
		fmt.Println(err.Error())
	}
        listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
        for {
		// AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn) // 这里是阻塞的，同时支持多个的话，需要用到go
	}
}
