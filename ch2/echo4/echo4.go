// echo4 输出其命令行参数package main
package main

import (
	"flag"
	"fmt"
	"strings"
)

// flag参数：标志、默认值、help信息，返回值是指针
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// 需要调用Parse()进行参数解析
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
