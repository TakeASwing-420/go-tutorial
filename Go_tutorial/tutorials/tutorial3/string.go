package main

import (
	"fmt"
	"os"
	"unicode/utf8"
	"strconv"
	 "reflect"
	"bufio"
	"strings"
)

func main() {
	var greeting string = "Hello, World 😊!"
	fmt.Println(greeting)
	var just string = "Hello, World!"
	fmt.Println(just)

	length := utf8.RuneCountInString(greeting)
	fmt.Printf("Length of greeting: %d\nLength of second one is: %d\n", length, len(just))

	var myrune rune = '5'
	myRune2 := '😊'
	fmt.Println(myrune, myRune2) // 53 128522
	a:=strconv.Itoa(90) //Integer to string
	fmt.Println(a)
	fmt.Printf("Type of text: %v\n", reflect.TypeOf(a))
	val,ok := strconv.Atoi("345");
	if  ok == nil {
		fmt.Printf("Type of text: %v\n", reflect.TypeOf(val))
	}else {
		fmt.Println("Value is not a string")
	}
	list := Take_input_list()
	if list != nil {
			fmt.Println("You entered:", list)
	}
}

func Take_input_list() []int {
		// Create a new reader to read input from the standard input (console)
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter a list of integers separated by space:")

		// Read the input string
		input, err := reader.ReadString('\n')
		if err != nil {
				fmt.Println("Error reading input:", err)
				return nil
		}

		// Trim the newline character from the input
		input = strings.TrimSpace(input)

		// Split the input string by spaces to get individual elements
		strElements := strings.Split(input, " ")

		// Create a slice to store the integers
		var intSlice []int

		// Convert each element to an integer and add it to the slice
		for _, str := range strElements {
				num, err := strconv.Atoi(str)
				if err != nil {
						fmt.Println("Error converting to integer:", err)
						return nil
				}
				intSlice = append(intSlice, num)
		}

		return intSlice
}