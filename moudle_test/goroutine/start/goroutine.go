package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine 是 go 语言中轻量级的线程
func routine()  {
	i := 0;
	for {
		i++
		fmt.Println("goroutine1:", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 启动一个 goroutine
  go routine()
  // 防止主线程退出 执行第一个 goroutine
  time.Sleep(3 * time.Second)
	fmt.Println("=======================")
	// 创建匿名函数
	go func() {
		defer fmt.Println("A ing...")
		func ()  {
			defer fmt.Println("B ing...")
			runtime.Goexit() // 终止当前 goroutine
		}()
		fmt.Println("A end...")
	}()
	// 防止主线程退出
	time.Sleep(1 * time.Second)
		
}