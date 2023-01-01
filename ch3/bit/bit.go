package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5

	fmt.Printf("%08b\n", x)
}
