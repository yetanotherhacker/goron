package main

import (
	"fmt"
	"log"
)

func basisCodeDistance(basis []uint, threshold uint) (isValid bool, distance uint, messageStatus string) {
	//basis of vectors represented as uints since the goal is small-dimensional search
	if len(basis) < 2 {
		return false, 0, "Need a multi-element array for a valid basis"
	}
	//start with maximum possible and descend
	var minDistance uint = (1 << 63) - 1
	for i, vectorA := range basis {
		subBasis := basis[i+1:]
		for _, vectorB := range subBasis {
			vectorDistance := uint(hamming(vectorA, vectorB))
			if minDistance > vectorDistance {
				minDistance = vectorDistance
				if threshold > minDistance {
					return false, 0, "Basis distance is lower than threshold."
				}
			}
		}
	}
	return true, minDistance, "Linear distance found."
}

func errorHandler(isValid bool, messageStatus string) {
	if !isValid {
		log.Fatal("Error: ", messageStatus)
	}
}

func findNextBasisVector(subBasis []uint, startPosition uint, distance uint, maxElement uint) (isValid bool, vector uint, messageStatus string) {
	//for incremental search of n -> n+1 vector space with desired minimum distance between codewords
	for candidate := startPosition; maxElement > candidate; candidate++ {
		var status, _, _ = basisCodeDistance(append(subBasis, candidate), distance)
		if !status {
			break
		} else {
			return true, candidate, "New basis vector found."
		}
	}
	return false, 0, "No valid basis vector found."
}

func hamming(x, y uint) uint {
	var (
		diff uint = x ^ y
		sum  uint = 0
	)
	for diff > 0 {
		sum += diff % 2
		diff = diff >> 1
	}
	return sum
}

func main() {
	example := []uint{1, 7, 255}
	var status, distance, msg = basisCodeDistance(example, 0)
	errorHandler(status, msg)
	fmt.Println(msg, distance)
}
