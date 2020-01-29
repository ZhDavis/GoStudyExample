package main

import (
    "fmt"
    _ "./b"
)
func init() {
    fmt.Println("c.go init1")
}
func init() {
    fmt.Println("c.go init2")
}

func main() {
    fmt.Println("main start!")
    fmt.Println("main end!")
}
