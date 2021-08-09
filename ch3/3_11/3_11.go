package main

import (
	"fmt"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma_sign_float(s string, sign bool) string {
	// 支持浮点数和正负数符号
	i := strings.Index(s, ".")
	if i == -1 {
		return comma(s)
	} else {
		return comma(s[:i]) + "." + comma(s[i+1:])
	}
}

func main() {
	fmt.Println(comma_sign_float("1234.56789", false))
	// s := "hello world"
	// fmt.Println(s[0])
	// fmt.Println(len(s))

	// a := [...]int{1, 2, 3, 4} // 长度为4的数组
	// b := a[0:2]               // 切片，len为2，cap为4
	// fmt.Println(len(b), cap(b))
	// fmt.Println(b[1])
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1))

	m := map[int]int{
		1: 2,
		3: 4,
	}
	fmt.Printf("%x", &m)

	type Point struct{ x, Y int }

	// # 法一
	p1 := Point{1, 2}

	// # 法二
	var p2 = Point{
		x: 1,
		Y: 2,
	}
	p3 := Point{x: 1, Y: 2}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}

func add(a ...int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
