package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func exercise(a [32]byte, b [32]byte) int {
	// 判断两个哈希码中不同的bit数，一个byte是8bit
	var count int
	for i := 0; i < 32; i++ {
		count += int(pc[a[i]^b[i]])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(exercise(c1, c2))
	fmt.Println(exercise(c1, c1))
	fmt.Printf("%x\n%x\n", c1, c2)
}
