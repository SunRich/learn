package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var s =[]int64{1,2,3}
	setSlice(s)
	fmt.Println(s)
	p := (*struct {

		array uintptr

		len   int

		cap   int

	})(unsafe.Pointer(&s))
	fmt.Printf("main output: %+v\n", p)
}

func setSlice(s []int64)  {

	s[1]=5
	s= append(s, 5)
	p := (*struct {

		array uintptr

		len   int

		cap   int

	})(unsafe.Pointer(&s))
	fmt.Printf("output: %+v\n", p)
}
