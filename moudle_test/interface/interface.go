package main

import "fmt"

// 接口定义
type Animal interface {
	Speak()
	Sleep()
}

func Speak() {
	fmt.Println("hello animal")
}

// 接口实现1
type Dog struct {
	Name string
	Age  int
}

func (d *Dog) Speak() {
	fmt.Println(d.Name + " speak")
}
func (d *Dog) Sleep() {
	fmt.Println(d.Name + " sleep")
}

// 接口实现2
type Cat struct {
	Name string
	Age  int
}

func (c *Cat) Speak() {
	fmt.Println(c.Name + " speak")
}
func (c *Cat) Sleep() {
	fmt.Println(c.Name + " sleep")
}

func main() {
	var animal Animal
	// 向下转型
	animal = &Dog{Name: "dog", Age: 1}
	animal.Speak()
	animal.Sleep()

	var cat Cat
	cat.Name = "cat"
	cat.Age = 2
	// 类型转换
	animal = &cat
	animal.Speak()
	animal.Sleep()
}
