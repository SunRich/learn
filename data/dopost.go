package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration)  {
	st:=time.Now()
	defer fmt.Println("print",time.Now().Sub(st))
	defer func() {
		fmt.Println("func",time.Now().Sub(st))
	}()
	time.Sleep(d)
}

func main()  {
	sleep(time.Second)
}
