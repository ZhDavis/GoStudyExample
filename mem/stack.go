package main 

func main() {
    fun()
}
//go:noinline
func fun() int {
    i := 10
    return i
}
