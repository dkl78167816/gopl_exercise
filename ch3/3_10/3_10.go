// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
)

// without recursion
func comma(s string) string {
	res := ""
	n := 3
	for ; n <= len(s); n += 3 {
		res += s[n-3:n] + ","
	}
	if n-3 != len(s) {
		res += s[n-3:]
	}

	if res[len(res)-1] == ","[0] {
		res = res[:len(res)-1]
	}

	// 解码为uft-8再比较
	// if r, n := utf8.DecodeLastRuneInString(res); r == rune(","[0]) {
	// 	res = res[:len(res)-n]
	// }
	return res
}

//  without recursion and use bytes.Buffer
func comma_2(s string) string {
	var res bytes.Buffer
	for i := 0; i < len(s)-1; i += 1 {
		res.WriteByte(s[i])
		if (i+1)%3 == 0 {
			res.WriteByte(',')
		}
	}
	res.WriteByte(s[len(s)-1])
	return res.String()
}

// func task12(a string, b string) bool {
// 计数
// 	count := make(map bytes [], 10)

// 	for c := range a {

// 	}
// }

func main() {
	fmt.Println(comma("123456"))
	fmt.Println(comma_2("1234567"))
	// fmt.Println(task10_2("12345678"))
	// fmt.Println(task11("1234567", false))
	// fmt.Println(task11("1234.56789", false))
}
