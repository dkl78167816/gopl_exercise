package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Implement with iteration
func PopCount(x uint64) int {
	var res int
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

func main() {
	fmt.Println(PopCount(7))
	fmt.Println(PopCount(16))
}
