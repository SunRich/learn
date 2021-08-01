package main

import (

	"fmt"

	"unsafe"

)


/**
struct String

{

        byte*   str;

        intgo   len;

};
 */
func main() {

	var str string = "hi, 陈一回~"

	p := (*struct {

		str uintptr

		len int

	})(unsafe.Pointer(&str))



	fmt.Printf("%+v\n", p)

}
