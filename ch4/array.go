package main

import "fmt"

func main() {
	// test array
	a := [...]int{1, 2, 3}
	fmt.Println(a)

	var b [4]int
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", b[:2])
	fmt.Println(b[:])
	fmt.Println(cap(b))

	var c [2]int = [2]int{1}
	fmt.Println(c)

	var d = [...]string{1: "1", 2: "3"}
	fmt.Printf("%T\n", d)
}
