package main

import "fmt"
// 任意类型参数
func judge(arg interface{}) {
	// 类型断言
	value, ok := arg.(int)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("int类型不匹配")
	}
	fmt.Printf("%T\n", arg)
}
func main() {
	a := 1
	b := "hello"
	judge(a)
	fmt.Println("--------")
	judge(b)
}
