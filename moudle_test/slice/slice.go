package main

import (
	"fmt"
)
// 切片是引用类型，切片的赋值操作会共享内存
func print(arr []int) {
	fmt.Println(len(arr), cap(arr), arr)
}
func main() {
	// 初始化
	arr := make([]int, 3)
	fmt.Println(len(arr), cap(arr), arr)
	// 追加
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	print(arr)
	// 截取
	arr1 := arr[3:]
	print(arr1)
	// 判断是否为空
	var nums []int
	if nums == nil {
		fmt.Println("nums is nil")

	} else {
		fmt.Println("nums is not nil")
	}
}
