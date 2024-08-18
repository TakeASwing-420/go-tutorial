package main

import "fmt"

type Number interface {
	int | int64 | float32 | float64
}

type Stack[T Number] struct {
	values []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(num T) {
	s.values = append(s.values, num)
}

// Pop removes and returns the element from the top of the stack
func (s *Stack[T]) Pop() T {
	if len(s.values) == 0 {
		panic("stack is empty")
	}
	index := len(s.values) - 1
	value := s.values[index]
	s.values = s.values[:index]
	return value
}

// Peek returns the element from the top of the stack without removing it
func (s *Stack[T]) Peek() T {
	if len(s.values) == 0 {
		panic("stack is empty")
	}
	return s.values[len(s.values)-1]
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("Top of intStack:", intStack.Peek())
	fmt.Println("Popped element from intStack:", intStack.Pop())
	fmt.Println("Top of intStack after pop:", intStack.Peek())

	floatStack := Stack[float64]{}
	floatStack.Push(3.14)
	floatStack.Push(6.28)

	fmt.Println("Top of floatStack:", floatStack.Peek())
	fmt.Println("Popped element from floatStack:", floatStack.Pop())
	fmt.Println("Top of floatStack after pop:", floatStack.Peek())
}
