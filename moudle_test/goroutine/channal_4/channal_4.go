package main

import (
	"fmt"
)

// select 语句
func fun(ch, exit chan int) {
	i := 0
	for {
		select {
		case ch <- i:
			i++
		case <-exit:
			return
		default:
			// 如果没有数据可读，则跳过
			continue
		}
	}
}
func main() {
	ch := make(chan int)
	exit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		exit <- 0
	}()
	fun(ch, exit)
}
