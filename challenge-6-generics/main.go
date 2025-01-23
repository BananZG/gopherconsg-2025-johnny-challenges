// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

func print[T string | int | bool](val T) {
	fmt.Println(val)
}

// Part 2 sum function refactoring
type numeric interface {
	constraints.Integer | constraints.Float
}

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	print("Hello")
	print(42)
	print(true)
	fmt.Println(sum(1, 2, 3))
}
