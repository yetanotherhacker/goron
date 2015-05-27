package main

import "fmt"

func hamming(x, y uint) uint {
	var (
		diff uint	= x^y
		sum	uint	= 0
	)
	for diff > 0 {
		sum += diff % 2
		diff = diff >> 1
	}
	return sum
}

func codeDistance(basis []uint) (string, int) {
	if len(basis) < 2 {
		return "Need more than one value for a valid basis", -1
	}
	var minDistance int = (1 << 63) - 1
	for i, va := range basis {
		subBasis := basis[i+1:]
		for _, vb := range subBasis {
			vectorDistance := int(hamming(va, vb))
			if minDistance > vectorDistance {
				minDistance = vectorDistance
			}
		}
	}
	return "Linear Code Distance:", minDistance
}

func main() {
	example := []uint{1, 7, 255}
	fmt.Println(codeDistance(example))
}
