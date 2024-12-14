package main

import (
	"fmt"
	"sync"
)

func main() {
  var m = sync.Map{}
	go func() {
		for {
			m.Store("a", 8)
		}
	}()

	go func() {
		for  {
			val,ok := m.Load("a")
			fmt.Println(val,ok)
		}
	}()
	select{}
}