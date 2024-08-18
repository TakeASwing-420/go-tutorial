package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i%3 == 0 {
			fmt.Printf("%s", "foo")
		} else if i%5 == 0 {
			fmt.Printf("%s", "vek")
		}

		if x := 90; i%12 == 0 && x-i < 67 {
			fmt.Println(x)
			break
		}
	}
	Hello() //How to build a package utility over multiple files is described here smoothly
   //'run following'
	//-->go build .
	//-->.\tutorial
}
