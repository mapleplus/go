package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
func (u User) Call()  {
	fmt.Println("user is called....")
}

func main() {
  u := User{1, "stu01", 20}
  DoFieldAndMethod(u)
}

func DoFieldAndMethod(arg interface{})  {
	//获取结构体类型和名称
	argType := reflect.TypeOf(arg)
	fmt.Println("type:", argType.Name())
	argValue := reflect.ValueOf(arg)
	fmt.Println("value: ", argValue)
	// 获取结构体内部字段
	for i := 0; i < argType.NumField(); i++{
		//获取字段类型
		field := argType.Field(i)
		// 获取字段值
		value := argValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// 获取结构体内部方法
	for i := 0; i < argType.NumMethod(); i++{
		m := argType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}