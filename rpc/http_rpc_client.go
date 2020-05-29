package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var a int = 10
        var outStr string
        fmt.Println("Client:send:",a) 
	err = client.Call("HttpRpcServer.Print", &a, &outStr)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("CLient revieve:",a,outStr)
}
