package main

import "fmt"

func main() {
	var isactive, flag bool // You can declare multiple variables side by side
	fmt.Scanf("%t", &isactive)
	fmt.Printf("%t", isactive)
	fmt.Println()
	fmt.Scan(&flag)
	fmt.Printf("You entered: %t\n", flag)
	var a uint8 = 67
	var b uint32 = 6416
	var c int64 = -53689992992
	fmt.Println(a, b, c)
	var x float32 = 3.14
	var y float64
	var z float32
	fmt.Scan(&y, &z)
	fmt.Println(x, y, z)
	const hello string = "Hello World"
	fmt.Println(hello)
	var z1 complex64 = complex(1.1, 2.2)
	var z2 complex128 = complex(3.3, 4.4)
	fmt.Println(z1, z2) // Output: (1.1+2.2i) (3.3+4.4i)
	Hello(45)
}

func Hello(a uint64) {

	var c uint64 = a //type declared
	var b = a        //type inferred
	d := a           //short hand declarartion
	fmt.Println(b, c, d)
}


//Mutable Types in Golang are Array, Map, Slice, Channel