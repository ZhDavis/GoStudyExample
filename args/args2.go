package main
import (
    "fmt"
    "flag"
)
var b = flag.Bool("b", false, "bool类型参数")
var s = flag.String("s", "", "string类型参数")
var t = flag.Int("t", 0, "int类型参数")

func argsfunc() {
    flag.Parse()
    fmt.Println("-b:", *b)
    fmt.Println("-s:", *s)
    fmt.Println("-t:", *t)
    fmt.Println("其他参数：", flag.Args())
    return
}

func main() {
    argsfunc()
    return
}
