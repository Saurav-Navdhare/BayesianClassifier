package main

import (
	"BayesianClassifier/model"
	"BayesianClassifier/utils"
	"fmt"
	"log"
	"os"
)

func trainModel(filePath string) {
	// Load the dataset
	data, err := utils.LoadData(filePath)
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	headers := data[0]
	data = data[1:]
	processedData := utils.BinaryLabelling(data)

	// convert to map[string][]int
	df := utils.ConvertToDF(processedData, headers)

	trainSet, testSet := model.TrainTestSplit(df, 0.2)

	// Calculate probabilities for training data
	classProbabilities, featureProbabilities := model.CalculateProbabilities(trainSet)

	// Evaluate the model on test set
	accuracy := model.Evaluate(testSet, classProbabilities, featureProbabilities)

	newModel := model.Model{
		ClassProbabilities: classProbabilities,
		FeatureStats:       featureProbabilities,
	}

	if accuracy > 0.85 {
		// Save the model
		err := model.SaveModel("model.json", newModel)
		if err != nil {
			log.Fatalf("Failed to save model: %v", err)
		}
	}

	// Print results
	fmt.Printf("Model accuracy: %.2f%%\n", accuracy*100)
}

func init() {
	// Set the output to be more descriptive
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// check if we already have weigths of the model or not in model.json file
	if _, err := os.Stat("model.json"); os.IsNotExist(err) {
		fmt.Println("Model not found. Training the model...")
		trainModel("data/diabetes_data_upload.csv")
	} else {
		fmt.Println("Model found. Loading the model...")
		_, err := model.LoadModel("model.json")
		if err != nil {
			log.Fatalf("Failed to load model: %v", err)
		}
	}
}

func main() {
	fmt.Println("Done")
}
