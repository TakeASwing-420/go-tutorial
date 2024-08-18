package main

import "fmt"

var (
	numbers = []int{1, 2, 3, 4, 5} //global variable
)

func main() {
	var arr [5]int // Declare an array of 5 integers
	arr[0] = 10    // Assign a value to the first element
	arr[1] = 20    // Assign a value to the second element
	arr = [5]int{12, 45, 67, 89, 90}
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
	hello := [...]string{"Go", "Python"} //use ...  i.e.  ellipsis operator for defining arrays slices length without specifying
	fmt.Println(hello)
	fmt.Println("Before calling", numbers)
	Myslice()
	fmt.Println("After calling", numbers)
	Lets_delete(4)
	fmt.Println("After deleting at index 4", numbers)
	sum(45, numbers...)
	const row, col = 2, 3
	arr2D := [row][col]int{{3, 45, 89}, {879, 6, 89}}
	fmt.Println(arr2D)

	a := make([]int, 1, 2) //[map,slice,chan] only used to create an empty instance of that type
	//first declare length i.e. 1 and then capacity i.e. 2
	a = append(a, 5)
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5) //capacity becomes 4
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5)
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5)
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5) //capacity becomes 8
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5)
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5)
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))
	a = append(a, 5) //capacity becomes 16
	fmt.Println(a, "is of length", len(a), "and of capacity", cap(a))

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Print("hello: ")
	fmt.Println(slice1, slice2)
}

func sum(a uint64, nums ...int) uint64 { //ellipsis operator ... in use
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total is", total)
	return uint64(total) + a
}

func Myslice() {
	additionalNumbers := []int{90, 89, 90}
	numbers = append(numbers, additionalNumbers...)
}

func Lets_delete(delindex uint8) {
	numbers = append(numbers[:delindex], numbers[delindex+1:]...)
}
