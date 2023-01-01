package comma // import "github.com/AntiBargu/gopl/ch3/comma"

import (
	"bytes"
)

func Comma(s string) string {
	var buf bytes.Buffer

	n := len(s)

	if n == 0 {
		return buf.String()
	}

	i := n % 3

	if i == 0 {
		i = 3
	}

	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
