package b

import (
    "fmt"
    "../a"
)
func B() {
    fmt.Println("b() start!")
    a.A()
    fmt.Println("b() end!")
}
func init() {
    fmt.Println("b.go init1")
}
func init() {
    fmt.Println("b.go init2")
}
