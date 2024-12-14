package main
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Des string `json:"des"`
}

func main() {
	// 创建一个Person对象
	p := Person{
		Name: "Alice",
		Age:  25,
		Des: "Hello, World!",
	}

	// 将Person对象转换为JSON字符串
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
	}else {
		fmt.Println("JSON data:", string(jsonData))
	}
	
	var p2 Person
	// 将JSON字符串转换为Person对象
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
	}else {
		fmt.Println("Unmarshaled data:", p2)
	}
}