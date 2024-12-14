package main

import (
	"fmt"
	"time"
)

// 有缓冲通道 关闭特性
// 1. 关闭通道后，可以继续读取数据，直到通道内所有数据被读取完毕
// 2. 关闭通道后，不能再写入数据
func main() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// 关闭通道
		close(ch)
	}()
	time.Sleep(time.Second)
	for {
		if num, ok := <-ch; ok {
			fmt.Printf("receive: %d\n", num)
		}else {
			break
		}
	}
	// 等价于
	// for num := range ch {
	// 	fmt.Printf("receive: %d\n", num)
	// }
}
