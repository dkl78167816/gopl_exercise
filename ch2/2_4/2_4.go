package main

import "fmt"

var pc [256]byte

func init() {
	// 统计从1到256，每个数都有多少个1bit
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Using the displacement algorithm, 1 bit per test
func PopCount(x uint64) int {
	var res = 0
	for x > 0 {
		res += int(x & 1)
		x >>= 1
	}
	return res
}

func main() {
	fmt.Println(PopCount(3))
	fmt.Println(PopCount(4))
	fmt.Println(PopCount(5))
}
