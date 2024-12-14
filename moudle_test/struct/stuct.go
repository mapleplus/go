package main

import "fmt"

// 父类
type Person struct {
	Name string
	Age  int
}

// 父类方法
func (p *Person) Eat() {
	fmt.Println("eating...")
}
func (p *Person) Sleep() {
	fmt.Println("sleeping...")
}

// 子类
type Student struct {
	Person
	Score int
}

// 子类重写父类方法
func (s *Student) Eat() {
	fmt.Println("Student eating...")
}
func (s *Student) Sleep() {
	fmt.Println(s.Name + " sleeping...")
}

// 子类新增方法
func (s *Student) Study() {
	fmt.Println("studying...")
}

func main() {
	// 父类实例
	var person Person
	person.Name = "zhangsan"
	person.Age = 18
	fmt.Println(person)
	person.Eat()
	person.Sleep()

	fmt.Println("------------")
	// 子类实例
	var student Student
	student.Name = "lisi"
	student.Age = 20
	student.Score = 100
	fmt.Println(student)
	student.Eat()
	student.Sleep()
	student.Study()
}
