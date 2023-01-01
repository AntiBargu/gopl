package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		// 每次DecodeRuneIn5tring的调用都返回r（文字符号本身）和一个值（表示r按UTF—8编码所占用的字节数）。
		// 这个值用来更新下标i，定位字符串内的下一个文字符号。对于非ASCII文字符号，下标增量大于1。
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// Go的range循环也适用于字符串，按UTF-8隐式解码。
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%[2]d\n", i, r)
	}
}
