package main

import (
	"fmt"
	"reflect"
)

// 标签 类似Java中的注解
type Person struct {
	Name string `info:"name" doc:"姓名"`
	Age  int    `info:"age" doc:"年龄"`
}

func fun(p interface{}) {
	// 使用反射获取 Person 对象的元素
	ptype := reflect.TypeOf(p).Elem()
	for i := 0; i < ptype.NumField(); i++ {
		// 获取标签内容
		field1 := ptype.Field(i).Tag.Get("info")
		field2 := ptype.Field(i).Tag.Get("doc")
		fmt.Println("field1: " + field1 + " field2: " + field2)
	}
}
func main() {
	// 创建一个 Person 对象
	var p Person = Person{Name: "张三", Age: 18}
	fun(&p)
}
