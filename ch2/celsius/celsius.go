package main

import "fmt"

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	var c Celsius = 100.0

	fmt.Printf("%s\n", c)
}
