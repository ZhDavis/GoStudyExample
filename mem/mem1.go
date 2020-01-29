
package main
import "fmt"
func main() {
	a := fun()
        fmt.Println("&a:",&a,"a:",a)
}

//go:noinline
func fun() int {
	i := new(int)
        *i = 666
        fmt.Println("&i:",&i,"i:",i)
	return *i
}
