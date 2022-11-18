package main

import (
	"fmt"
	"sync"
)

var balance int = 0
var mu sync.Mutex
var w sync.WaitGroup

func Add(cash int) {
	mu.Lock()
	balance += cash
	mu.Unlock()
	defer w.Done()
}

func main() {
	w.Add(10)
	for i := 0; i < 10; i++ {
		go Add(i)
	}
	w.Wait()
	fmt.Println(balance)
}
