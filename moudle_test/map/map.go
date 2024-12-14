package main

import "fmt"

func main() {
	// 声明一个map
	map1 := make(map[string]int)
	// 添加键值对
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	print(map1)
	fmt.Println("=====================")
	// 删除键值对
	delete(map1, "two")
	// 修改
	map1["one"] = 11
	print(map1)
}
// 打印map
func print(mapTemp map[string]int) {
	for key, v := range mapTemp {
		fmt.Println("key:", key, "value:", v)
	}
}
