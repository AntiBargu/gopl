package comma1 // import "github.com/AntiBargu/gopl/ch3/comma1"

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	var buf bytes.Buffer

	if len(s) == 0 {
		return buf.String()
	}

	// 处理正负号
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}

	// 分成整数部和小数部，分别处理
	var intStr, floatStr string

	dotIdx := strings.Index(s, ".")
	if -1 == dotIdx {
		intStr = s
	} else {
		intStr, floatStr = s[:dotIdx], s[dotIdx:]
	}

	i := len(intStr) % 3
	if i == 0 {
		i = 3
	}

	buf.WriteString(intStr[:i])
	for ; i < len(intStr); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(intStr[i : i+3])
	}
	buf.WriteString(floatStr)

	return buf.String()
}
