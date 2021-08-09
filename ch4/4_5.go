package main

import (
	"fmt"
)

// 重写reverse函数
func exercise(s []string) []string {
	if len(s) == 0 {
		return s
	}
	head := 1
	last := s[0]
	for j := 1; j < len(s); j++ {
		if s[j] != last {
			s[head] = s[j]
			last = s[j]
			head += 1
		}
	}

	return s[:head]
}

// 使用切片+append，省略head
func exercise2(s []string) []string {
	if len(s) == 0 {
		return s
	}
	res := s[:1]
	last := s[0]
	for j := 1; j < len(s); j++ {
		if s[j] != last {
			res = append(res, s[j])
			last = s[j]
		}
	}
	return res
}

func main() {
	data := []string{"123", "456", "123", "123", "789"}
	fmt.Println(exercise(data))
	// 原数据已被更改
	fmt.Println(data)

	data = []string{"123", "456", "123", "123", "789"}
	fmt.Println(exercise2(data))
	fmt.Println(data)
}
