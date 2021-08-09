package main

import (
	"fmt"
)

// reverse reverses a slice of ints in place.
func reverseOrigin(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// replace silce with array pointer
func exercise(a *[6]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

func main() {
	data := [...]int{1, 2, 3, 4, 5, 6}
	reverseOrigin(data[:])
	fmt.Println(data)

	exercise(&data)
	fmt.Println(data)
}
