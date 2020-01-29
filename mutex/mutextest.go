package main  
import (  
    "fmt"
    "sync"
)

func main() {  
    var w sync.WaitGroup
    var m sync.Mutex
    var x int = 0
    for i := 0; i < 100; i++ {
        w.Add(1)        
        go func(){
            m.Lock()
            x = x + 1
            m.Unlock()
            w.Done()
        }()
    }
    w.Wait()
    fmt.Println("main value of x:", x)
}
