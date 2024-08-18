package main

type operation func(int, int) int //typedef in c

func add(a int, b int) int {
    return a + b
}

func applyOperation(a int, b int, op operation) int {
    return op(a, b)
}


