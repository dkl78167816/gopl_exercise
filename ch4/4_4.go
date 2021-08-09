package main

import (
	"fmt"
)

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// 通过一遍循环完成旋转，新建数组，计算新索引即可
// 这里是利用3次反转来实现（两遍循环）
func rotate(a []int, n int) {
	reverse(a)
	reverse(a[:n])
	reverse(a[n:])
}

func main() {
	data := []int{1, 2, 4, 5, 6, 7, 8}
	rotate(data, 3)
	fmt.Println(data)
}
