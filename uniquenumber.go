package main

// you can also use imports, for example:
import (
	"fmt"
	"sort"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution() int {
	var A = []int{1, 2, 4, 5, 6, 7}
	sort.Ints(A)

	var max int = 1
	for _, a := range A {
		if a == max {
			max++
		}
	}
	return max
}

func main() {
	result := Solution()
	fmt.Println(result)
}
