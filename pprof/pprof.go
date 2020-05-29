package main 
import (
    "fmt"
    "time"
    "sync"
    "net/http"
    _ "net/http/pprof"
)

var buf []byte
func Add() {
    tick := time.Tick(time.Second / 200)
    for range tick {
        buf = append(buf, make([]byte, 2*1024*1024)...)
    }
}
func main() {
    // 开启pprof，监听请求
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        ip := "0.0.0.0:6060"
        if err := http.ListenAndServe(ip, nil); err != nil {
            fmt.Printf("start pprof failed on %s\n", ip)
        }
    }()
    fmt.Println("continue~")
    Add()
    wg.Wait()
}
