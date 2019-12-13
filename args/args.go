package main

import (
    "fmt"
    "os"
    "strconv"
)

func argsfunc () {
    fmt.Println("argsfunc start")
    for idx, args := range os.Args {
        fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
    }
    fmt.Println("args func end")
    return
}

func main() {
    fmt.Println("main start")
    argsfunc()
    
    fmt.Println("main end")
}
