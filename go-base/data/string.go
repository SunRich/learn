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

	var str string = "hi, ιδΈε~"

	p := (*struct {

		str uintptr

		len int

	})(unsafe.Pointer(&str))



	fmt.Printf("%+v\n", p)

}
