package main

import (
    "fmt"
    "os"
    "strconv"
)

func main () {
    fmt.Println("main start,参数信息:")
    for idx, args := range os.Args {
        fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
    }
    fmt.Println("main end!")
    return 
}
