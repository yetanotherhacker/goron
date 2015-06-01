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

func codeDistance(basis []uint, threshold uint) (isValid bool, distance uint, messageStatus string) {
	//basis of vectors represented as uints since the goal is small-dimensional search
	if len(basis) < 2 {
		return false, 0, "Need a multi-element array for a valid basis"
	}
	var minDistance uint = (1 << 63) - 1
	for i, vectorA := range basis {
		subBasis := basis[i+1:]
		for _, vectorB := range subBasis {
			vectorDistance := uint(hamming(vectorA, vectorB))
			if minDistance > vectorDistance {
				minDistance = vectorDistance
				if (threshold > minDistance) {
					return false, 0, "Basis distance is lower than threshold."
				}
			}
		}
	}
	return true, minDistance, "Linear Code Distance found."
}

func main() {
	example := []uint{1, 7, 255}
	fmt.Println(codeDistance(example, 0))
}
