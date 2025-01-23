package main

import "fmt"

// 1. Declare a package-level variable val with inferred type string with an assigned value
var val = "global"

func main() {
	// 2. In main, use short variable declaration for a local variable val and assign it an integer
	val := 42

	// 3. Print out the type and value of the local val
	fmt.Printf("%T, local val = %v\n", val, val)

	// 4. Write a function to print out the type and value of the global val
	printGlobal()
	// 5. Write a function to update the global val
	updateGlobal()
	printGlobal()

	// 6. Print out the type and value of a dereference local val
	valPointer := &val
	fmt.Printf("%T, local &val = %v\n", valPointer, valPointer)
	// 7. Assign a value to the local val directly to the underlying memory location you - you will need to use a pointer dereference for this.
	*(valPointer) = 100

	// 8. Print out the local val to show you managed to update the underlying memory location.
	fmt.Printf("local val = %v\n", val)
}

func updateGlobal() {
	val = "updated global"
}

func printGlobal() {
	fmt.Printf("global val = %v\n", val)
}
