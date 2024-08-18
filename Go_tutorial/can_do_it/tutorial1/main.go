package main

import (
		"bufio"
		"fmt" //formatted input output full form
		"os"
)

func main() {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter some text: ")
		input, err := reader.ReadBytes('\n')
		if err != nil {
				fmt.Println("Error reading input:", err)
				return
		}

		fmt.Println("You entered (as bytes):", input)
		fmt.Println("You entered (as string):", string(input))
}
