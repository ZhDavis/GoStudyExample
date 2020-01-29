package a

import "fmt"

func A() {
    fmt.Println("a() start!")
}
func init() {
    fmt.Println("a.go init1")
}
func init() {
    fmt.Println("a.go init2")
}
