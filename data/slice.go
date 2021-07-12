package main

/*
struct  Slice

{               // must not move anything

byte*   array;      // actual data

uintgo  len;        // number of elements

uintgo  cap;        // allocated number of elements

};*/


import (

	"fmt"

	"unsafe"

)



func main() {

	var slice []int32 = make([]int32, 5, 10)

	p := (*struct {

		array uintptr

		len   int

		cap   int

	})(unsafe.Pointer(&slice))
	fmt.Printf("output: %+v\n", p)
	slice=slice[0:1]

	fmt.Printf("output: %+v\n", p)

}
