package model

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

type Model struct {
	ClassProbabilities map[int]float64
	FeatureStats       map[string]map[int]map[string]float64
}

// SaveModel function to save the model parameters to a file
func SaveModel(filename string, model Model) error {
	for feature := range model.FeatureStats {
		for class := range model.FeatureStats[feature] {
			if math.IsNaN(model.FeatureStats[feature][class]["mean"]) {
				model.FeatureStats[feature][class]["mean"] = 0
			}
			if math.IsNaN(model.FeatureStats[feature][class]["variance"]) {
				model.FeatureStats[feature][class]["variance"] = 0
			}
		}
	}

	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func LoadModel(filename string) (Model, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Model{}, err
	}

	var model Model
	if err := json.Unmarshal(data, &model); err != nil {
		return Model{}, err
	}

	return model, nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func CalculateProbabilities(trainSet map[string][]int) (map[int]float64, map[string]map[int]map[string]float64) {
	classCounts := make(map[int]int)
	featureStats := make(map[string]map[int]map[string]float64)
	classProbabilities := make(map[int]float64)

	// Initialize maps
	for feature := range trainSet {
		if feature == "class" {
			continue
		}
		featureStats[feature] = make(map[int]map[string]float64)
		for class := range trainSet["class"] {
			featureStats[feature][class] = make(map[string]float64)
		}
	}

	// Calculate class counts and sums
	for i := 0; i < len(trainSet["class"]); i++ {
		class := trainSet["class"][i]
		classCounts[class]++

		for feature, values := range trainSet {
			if feature == "class" {
				continue
			}
			value := float64(values[i])
			featureStats[feature][class]["sum"] += value
			featureStats[feature][class]["sumSq"] += value * value
		}
	}

	total := len(trainSet["class"])
	for class, count := range classCounts {
		classProbabilities[class] = float64(count) / float64(total)
	}

	// Use concurrency to calculate mean and variance
	var wg sync.WaitGroup
	for feature := range featureStats {
		for class := range featureStats[feature] {
			wg.Add(1)
			go func(feature string, class int) {
				defer wg.Done()
				count := classCounts[class]
				if count > 0 {
					mean := featureStats[feature][class]["sum"] / float64(count)
					variance := (featureStats[feature][class]["sumSq"] / float64(count)) - (mean * mean)
					featureStats[feature][class]["mean"] = mean
					featureStats[feature][class]["variance"] = variance
					fmt.Printf("%sClass %d, Feature %s: Mean = %.2f, Variance = %.2f%s\n", Blue, class, feature, mean, variance, Reset)
				} else {
					featureStats[feature][class]["mean"] = math.NaN()
					featureStats[feature][class]["variance"] = math.NaN()
				}
			}(feature, class)
		}
	}
	wg.Wait()
	newModel := Model{
		ClassProbabilities: classProbabilities,
		FeatureStats:       featureStats,
	}
	return newModel.ClassProbabilities, newModel.FeatureStats
}

func gaussianProbability(x, mean, variance float64) float64 {
	exponent := math.Exp(-((x - mean) * (x - mean)) / (2 * variance))
	return (1 / (math.Sqrt(2 * math.Pi * variance))) * exponent
}

func Predict(sample map[string]int, classProbabilities map[int]float64, featureStats map[string]map[int]map[string]float64) int {
	bestClass := -1
	bestProb := -1.0

	for class, classProb := range classProbabilities {
		prob := classProb
		for feature, value := range sample {
			mean := featureStats[feature][class]["mean"]
			variance := featureStats[feature][class]["variance"]
			prob *= gaussianProbability(float64(value), mean, variance)
		}
		fmt.Printf("%sClass %d: Probability = %.6f%s\n", Yellow, class, prob, Reset)
		if prob > bestProb {
			bestClass = class
			bestProb = prob
		}
	}
	fmt.Printf("%sPredicted Class: %d%s\n", Green, bestClass, Reset)
	return bestClass
}

func Evaluate(testSet map[string][]int, classProbabilities map[int]float64, featureStats map[string]map[int]map[string]float64) float64 {
	truePositives, trueNegatives, falsePositives, falseNegatives := 0, 0, 0, 0
	total := len(testSet["class"])

	for i := 0; i < total; i++ {
		trueClass := testSet["class"][i]
		features := make(map[string]int)
		for feature, values := range testSet {
			if feature != "class" {
				features[feature] = values[i]
			}
		}

		predictedClass := Predict(features, classProbabilities, featureStats)

		if predictedClass == 1 && trueClass == 1 {
			truePositives++
		} else if predictedClass == 0 && trueClass == 0 {
			trueNegatives++
		} else if predictedClass == 1 && trueClass == 0 {
			falsePositives++
		} else if predictedClass == 0 && trueClass == 1 {
			falseNegatives++
		}
	}

	accuracy := float64(truePositives+trueNegatives) / float64(total)
	precision := float64(truePositives) / float64(truePositives+falsePositives)
	recall := float64(truePositives) / float64(truePositives+falseNegatives)
	f1Score := 2 * (precision * recall) / (precision + recall)

	fmt.Println("Confusion Matrix:")
	fmt.Printf("%sTP: %d%s | %sFP: %d%s\n", Green, truePositives, Reset, Red, falsePositives, Reset)
	fmt.Printf("%sFN: %d%s | %sTN: %d%s\n", Red, falseNegatives, Reset, Green, trueNegatives, Reset)
	fmt.Printf("%sAccuracy: %.2f%s\n", Cyan, accuracy, Reset)
	fmt.Printf("%sPrecision: %.2f%s\n", Yellow, precision, Reset)
	fmt.Printf("%sRecall: %.2f%s\n", Blue, recall, Reset)
	fmt.Printf("%sF1 Score: %.2f%s\n", Magenta, f1Score, Reset)

	return accuracy
}
