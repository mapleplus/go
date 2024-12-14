package main

import (
	"fmt"
	"time"
)

// 有缓冲通道
func main() {
	ch := make(chan int, 3)
	fmt.Printf("len:%d cap:%d\n", len(ch), cap(ch))
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("send: %d len: %d cap: %d\n",i, len(ch), cap(ch))
		}
	}()
	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("receive:", num)
	}
}
