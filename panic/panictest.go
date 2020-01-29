
package main

import (
	"fmt"
)

func testPanic() {
	fmt.Println("testPanic start!")
	defer func() {
		fmt.Println("testPanic defer!") // 1.

	}()
	panic("trigger panic!")       // 2.
	fmt.Println("testPanic end!") // 3.
}
func main() {
	fmt.Println("main start!")
	defer func() {
		fmt.Println("main defer!")
		if r := recover(); r != nil { // 4.
			fmt.Println("recovered from ", r)
		}
		fmt.Println("main exit!")
	}()
	go testPanic()           // 5.
	fmt.Println("main end!") // 6.
}

