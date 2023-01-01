package main

import "fmt"

func main() {
	s := "你好，世界"
	// 按16进制打印UTF-8编码
	fmt.Printf("% x\n", s)

	// 当[]rune转换作用于UTF—8编码的字符串时，返回该字符串的Unicode码点序列
	r := []rune(s)
	fmt.Printf("% x\n", r)

	// 如果把rune slice转换成一个字符申，它会输出各个文字符号的UTF—8编码拼接结果
	fmt.Println(string(r))

	// 若将一个整数值转换成字符串，其值按rune类型解读，并且string()转换产生代表该文字符号值的UTF-8码
	fmt.Println(string(0x4eac)) // "京"
}
