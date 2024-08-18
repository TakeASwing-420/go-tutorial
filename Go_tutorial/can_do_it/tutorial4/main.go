package main

import (
	"fmt"
	"go_tutorials/Go_tutorial/can_do_it/tutorial4/helper"
)

func main() {
	a, b := 56, 80
	x,_,_:= calc.Average(a, b) //unnamed variables
	
	fmt.Println(x) // Use helper package to access the Average function
}
