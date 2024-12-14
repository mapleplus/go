package main

import (
	"fmt"
	"sync"
)
var lock sync.Mutex
var wg sync.WaitGroup
var count int
func add()  {
	lock.Lock()
	defer lock.Unlock()
	for i := 0; i < 100; i++ {
		count++
	}
	wg.Done()
}
func del()  {
	lock.Lock()
	defer lock.Unlock()
	for i := 0; i < 100; i++ {
		count--
	}
	wg.Done()
}
func main() {
	wg.Add(2)
  go add()
	go del()
	wg.Wait()
	fmt.Println(count)
}