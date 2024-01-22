package main

import (
	"fmt"
	"slices"
)

func main() {

	names := []string{"john", "wick", "muhammad"}
	values := []int{10, 20, 100}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "john"))
	fmt.Println(slices.Index(names, "wick"))
	fmt.Println(slices.Index(names, "muhammad"))
}
