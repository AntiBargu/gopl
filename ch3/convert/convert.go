package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(strconv.Itoa(x), y)

	a, _ := strconv.Atoi("123")             // x是一个整数
	b, _ := strconv.ParseInt("123", 10, 64) // 十进制，最长为64位

	fmt.Println(a, b)
}
