package main

import (
	"fmt"
	"sync"
)

type Object struct {
	// You can define any fields you need for your object
	Value int
}

func bolvai() {
	// Initialize a new sync.Pool
	pool := sync.Pool{
		New: func() any {
			// This function is called when a new object is needed
			return &Object{Value : -1}
		},
	}

	// Put some objects into the pool
	for i := 0; i < 5; i++ {
		obj := &Object{Value: i}
		fmt.Println(obj)
		pool.Put(obj)
	}

	// Retrieve objects from the pool
	for i := 0; i < 5; i++ {
		obj := pool.Get().(*Object)
		fmt.Printf("Object %d from pool with value: %d\n", i, obj.Value)}
		// Do some processing with the object...
		// After use, put the object back into the pool
	

	fmt.Println()

	// Objects can be reused from the pool
	for i := 5; i < 10; i++ {
		obj := pool.Get().(*Object)
		fmt.Printf("Object %d from pool with value: %d\n", i, obj.Value)
		obj.Value += 1
		pool.Put(obj)
	}

	fmt.Println()
	// Retrieve and print objects from the pool until a new (empty) object is encountered
	for { //Currently pool has only one object
		obj := pool.Get().(*Object)
		if obj.Value == -1 {
			fmt.Println("Encountered an empty object, stopping the loop.")
			break
		} else {
			fmt.Printf("Object from pool with value: %d\n", obj.Value)
		}
	}

}
