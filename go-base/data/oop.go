package main

import "fmt"

func deferPrint(i int)  {
	fmt.Println(i)
}
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
