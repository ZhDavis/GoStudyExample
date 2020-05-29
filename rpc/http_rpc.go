package  main

import (
// 	"errors"
	"fmt"
        "strconv"
	"net/http"
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
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
