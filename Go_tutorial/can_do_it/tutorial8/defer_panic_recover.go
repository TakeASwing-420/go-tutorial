package main

import "fmt"

func eth() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
