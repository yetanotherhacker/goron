package main

import "fmt"

func hamming(x, y int) int {
	var (
		diff int	= x^y
		sum	int		= 0
	)
	for diff > 0 {
		sum += diff % 2
		diff = diff >> 1
	}
	return sum
}

func main() {
	fmt.Println(hamming(1,7))
}
