package main

import (
	"fmt"
	"time"
)

func readblock() {
	fmt.Println("Block Start!")
	c1 := make(chan int, 3) // c1 的容量是3，当写入的数据大于3的时候，写协程就会被阻塞住

	go func() { // 1.goroutine 用来往channel写数据，写的数据要多于3个
		for i := 0; i < 5; i++ {
			c1 <- i
			fmt.Printlnf("Write data into channel:", i, time.Now())
			time.Sleep(2 * time.Second)
		}
	}()

	// 2. 主 goroutine 用来从channel读数据，读数据读的比较慢
	for i := 0; i < 5; i++ {
		tmp := <-c1
		fmt.Println("Read data from channel:", tmp, time.Now())
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Block exit!")
}
func writeblock() {
	fmt.Println("Block Start!")
	c1 := make(chan int, 3) // c1 的容量是3，当写入的数据大于3的时候，写协程就会被阻塞住

	go func() { // 1.goroutine 用来往channel写数据，写的数据要多于3个
		for i := 0; i < 5; i++ {
			c1 <- i
			fmt.Println("Write data into channel:", i, time.Now())
		}
	}()

	// 2. 主 goroutine 用来从channel读数据，读数据读的比较慢
	for i := 0; i < 5; i++ {
		tmp := <-c1
		fmt.Println("Read data from channel:", tmp, time.Now())
		time.Sleep(2 * time.Second)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Block exit!")
}

func noblock() {
	fmt.Println("Main Start!")
	c1 := make(chan int, 10)

	go func() { // 1.goroutine 用来往channel写数据
		for i := 0; i < 5; i++ {
			c1 <- i
			fmt.Println("Write data into channel:", i, time.Now())
		}
	}()

	// 2. 主 goroutine 用来从channel读数据
	for i := 0; i < 5; i++ {
		tmp := <-c1
		fmt.Println("Read data from channel:", tmp, time.Now())
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Main exit!")
}

func main() {
	readblock()
	return
}
