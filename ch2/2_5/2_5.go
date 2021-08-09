package main

import "fmt"

var pc [256]byte

func init() {
	// 统计从1到256，每个数都有多少个1bit
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Using x&(x-1)
func PopCount(x uint64) int {
	var res = 0
	for x > 0 {
		x = x & (x - 1)
		res += 1
	}
	return res
}

func main() {
	fmt.Println(PopCount(8))
	fmt.Println(PopCount(10))
	fmt.Println(PopCount(15))
}
