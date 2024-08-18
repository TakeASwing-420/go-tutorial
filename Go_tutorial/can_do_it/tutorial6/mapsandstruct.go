package main

import (
	"fmt"
)

// Define the Person struct
type Person struct {
	Name string
	Age  int
}

// Declare a map from string (address) to Person struct
var people map[string]Person

// Initialize the map in an init function or in main
func init() {
	people = make(map[string]Person)
}

// Function to set a person's data
func setPerson(address, name string, age int) {
	people[address] = Person{Name: name, Age: age}
}
// Function to delete a person's data
func deletePerson(address string) bool {
		if _, exists := people[address]; exists {
				delete(people, address)
				return true
		}
		return false
}
// Function to get a person's data
func getPerson(address string) (Person, bool) {
	person, exists := people[address]
	return person, exists
}

func main() {
	address1 := "0x1234"
	address2 := "0x5678"

	setPerson(address1, "Alice", 30)
	setPerson(address2, "Bob", 25)

	person1, exists1 := getPerson(address1)
	if exists1 {
		fmt.Printf("Person at address %s: %+v\n", address1, person1)
	} else {
		fmt.Printf("No person found at address %s\n", address1)
	}

	person2, exists2 := getPerson(address2)
	if exists2 {
		fmt.Printf("Person at address %s: %+v\n", address2, person2)
	} else {
		fmt.Printf("No person found at address %s\n", address2)
	}

	// Trying to get a non-existent address
	person3, exists3 := getPerson("0x9999")
	if exists3 {
		fmt.Printf("Person at address %s: %+v\n", "0x9999", person3)
	} else {
		fmt.Printf("No person found at address %s\n", "0x9999")
	}
	 // Delete a person's data
		deleted := deletePerson(address1)
		if deleted {
				fmt.Printf("Person at address %s deleted.\n", address1)
		} else {
				fmt.Printf("No person found at address %s to delete.\n", address1)
		}

		// Try to get the deleted person's data
		person1, exists1 = getPerson(address1)
		if exists1 {
				fmt.Printf("Person at address %s: %+v\n", address1, person1)
		} else {
				fmt.Printf("No person found at address %s\n", address1)
		}
	}
}
