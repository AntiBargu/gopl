package main

import (
	"bufio"
	"fmt"
	"os"
)

type LineInfo struct {
	count int
	files []string
}

func main() {
	lineMap := make(map[string]*LineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lineMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, lineMap)
			f.Close()
		}
	}
	for line, n := range lineMap {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n", n.count, line)
		}
	}
}

func countLines(f *os.File, lineMap map[string]*LineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, exist := lineMap[line]; exist {
			lineMap[line].count++
			lineMap[line].files = append(lineMap[line].files, f.Name())
		} else {
			lineMap[line] = &LineInfo{1, []string{f.Name()}}
		}
	}
	// 注意: 忽略input.Err()中可能出现得错误
}
