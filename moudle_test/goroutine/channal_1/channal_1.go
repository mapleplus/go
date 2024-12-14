package main

import (
	"fmt"
	"time"
)

// 无缓冲通道
func main() {
  ch := make(chan int)
	fmt.Printf("len:%d cap:%d\n", len(ch), cap(ch))
  go func() {
    for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("send:", i)
			time.Sleep(time.Second)
		}
  }()
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("receive:", num)
	}
}