package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func half(b int) (int, bool) {
	helper2 := func() (ret int, success bool) {
		ret = b >> 1
		success = (b%2 == 0)
		return
	}
	return helper2()
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var num int
	fmt.Print("Enter the number: ")
	if reader.Scan() {
		input := reader.Text()
		var err error
		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			return
		}
	}
	halfValue, isEven := half(num)
	fmt.Printf("Half: %d, Is even? %t\n", halfValue, isEven)

}
