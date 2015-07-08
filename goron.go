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
			vectorDistance := uint(hammingDistance(vectorA, vectorB))
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

func bestVectorSearch(subBasis []uint, maximin uint, maxElement uint) (isValid bool, vector uint, distance uint, messageStatus string) {
	var (
		candidate     uint
		isSearchValid bool = true
		newDistance   uint
		startPosition uint = 0
	)
	for isSearchValid {
		isSearchValid, candidate, newDistance, _ = validVectorSearch(subBasis, startPosition, maximin, maxElement)
		if newDistance > maximin {
			maximin = newDistance
		}
		startPosition = candidate + 1
	}
	if startPosition > 0 {
		return true, startPosition, maximin, "New best-distance vector found"
	} else {
		return false, startPosition, maximin, "No valid basis vector found"
	}
}

func errorHandler(isValid bool, messageStatus string) {
	if !isValid {
		log.Fatal("Error: ", messageStatus)
	}
}

func hammingDistance(vectorA, vectorB uint) uint {
	var (
		diff uint = vectorA ^ vectorB
		sum  uint = 0
	)
	for diff > 0 {
		sum += diff % 2
		diff = diff >> 1
	}
	return sum
}

func validVectorSearch(subBasis []uint, startPosition uint, distance uint, maxElement uint) (isValid bool, vector uint, newDistance uint, messageStatus string) {
	//for incremental search of n -> n+1 vector space with desired minimum distance between codewords
	for candidate := startPosition; maxElement > candidate; candidate++ {
		var status, newDistance, _ = basisCodeDistance(append(subBasis, candidate), distance)
		if !status {
			continue
		} else {
			return true, candidate, newDistance, "New basis vector found."
		}
	}
	return false, 0, newDistance, "No valid basis vector found."
}

func main() {
	example := []uint{1, 14, 112}
	var status, distance, msg = basisCodeDistance(example, 0)
	errorHandler(status, msg)
	fmt.Println(msg, distance)
	status, newBasis, _, msg := validVectorSearch(example, 0, 4, 255)
	errorHandler(status, msg)
	fmt.Println(msg, newBasis)
}
