package main

import (
    "fmt"
    "sync"
    "time"
)

var m *sync.RWMutex
var num int = 0
func main() {
    m = new(sync.RWMutex)
    for i := 0; i < 5; i++ {
        go  write(i)
    }
    go  read(0)
    time.Sleep(10*time.Second)
}

func write(i int) {
    m.Lock()
    num = num + 1
    fmt.Println(i,"writing", num)
    fmt.Println(i,"writing done")
    m.Unlock()    
}

func read(i int) {
    fmt.Println(i,"reading start")
    m.RLock()
    fmt.Println(i,"reading",num)
    time.Sleep(1*time.Second)
    m.RUnlock()    
    fmt.Println(i,"reading done")
}

