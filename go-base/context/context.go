package main

import (
	"fmt"
	"sync"
)

type s struct {
	m     map[int]int
	mutex sync.Mutex
}

func main() {
	defer func() {
		fmt.Println("defer 1")
		if x := recover(); x != nil {
			// TODO fix panic
			fmt.Println(222,x)
		}
	}()


	defer func() {
		fmt.Println("defer 2")
		defer func() {
			fmt.Println("defer 3")

		}()
		panic("panic again and again")
	}()
	panic(222)
}

func panicDows()  {
	go func() {
		defer println("in goroutine")
		panic("1")
	}()
}
