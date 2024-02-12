package main

import (
	"fmt"
	"sort"
)

// variable should be declared at root level to be dscovered inside main are not accesible by other methods.
func main() {
	sayhi("Harsha")

	for _, v := range points {
		fmt.Print(v, "  ")

	}
	sort.Ints(points)
	fmt.Print(points)
	fmt.Print(hello)
}
