package model

import (
	"math/rand"
	"time"
)

func SplitInXY(df map[string][]int, output string) (map[string][]int, []int) {
	X := make(map[string][]int)
	Y := make([]int, len(df[output]))
	for key := range df {
		if key == output {
			Y = df[key]
		} else {
			X[key] = df[key]
		}
	}
	return X, Y
}

func TrainTestSplit(df map[string][]int, testRatio float64) (map[string][]int, map[string][]int) {
	rand.Seed(time.Now().UnixNano())
	n := len(df["class"])
	trainSize := int(float64(n) * (1 - testRatio))

	// Shuffle indices
	indices := rand.Perm(n)

	trainSet := make(map[string][]int)
	testSet := make(map[string][]int)

	// Allocate slices for each feature
	for feature := range df {
		trainSet[feature] = make([]int, trainSize)
		testSet[feature] = make([]int, n-trainSize)
	}

	// Fill in train and test sets
	for i, idx := range indices {
		for feature, values := range df {
			if i < trainSize {
				trainSet[feature][i] = values[idx]
			} else {
				testSet[feature][i-trainSize] = values[idx]
			}
		}
	}
	return trainSet, testSet
}
