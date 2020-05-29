package main

import (
	"fmt"
)

func ChannelCache() {
	fmt.Println("ChannelCache: start!")
	defer fmt.Println("ChannelCache: exit!")
	c := make(chan int, 10)
	defer close(c)
	go func() {
		c <- 2 + 3
	}()
	i := <-c
	fmt.Println("i:", i)
	return
}

func ChannelNoCache() {
	fmt.Println("ChannelNoCache: start!")
	defer fmt.Println("ChannelNoCache: exit!")
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 2 + 3
	}()
	i := <-c
	fmt.Println("i:", i)
	return
}

func main() {
	fmt.Println("Main start!")
	defer fmt.Println("Main exit!")

	ChannelNoCache()
	ChannelCache()
	return
}
