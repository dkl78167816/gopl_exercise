package main

import (
	"fmt"
	"unicode"
)

// 重写reverse函数，用于替换utf-8编码中[]byte类型的相邻空格
func exercise(s []byte) []byte {
	if len(s) == 0 {
		return s
	}

	cont := false
	res := s[:0]
	for _, v := range s {
		if unicode.IsSpace(rune(v)) {
			if !cont {
				cont = true
				res = append(res, v)
			}
			continue
		}
		res = append(res, v)
		cont = false
	}
	return res
}

func main() {
	data := []byte("你好，  世   界")
	fmt.Printf("%s\n", exercise(data))
	// 原数据已被更改
	fmt.Printf("%s", data)
}
