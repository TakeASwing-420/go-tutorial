package main

import "fmt"

func Hello() {
	var number uint = 3

	// switch number{
	// case 0:
	//     fmt.Println("Zero")
	// case 1:
	//     fmt.Println("One")
	// default:
	//     fmt.Println("Unknown :)")
	// }

	switch {
	case number == 0:
		fmt.Println("Zero")
	case number == 1:
		fmt.Println("One")
	default:
		fmt.Println("Unknown :P")
	}
}
