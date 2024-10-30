# Project Title

This project implements a machine learning model in Go, focusing on data preprocessing, training, and evaluation. The model handles binary classification tasks and includes functions for splitting data, calculating probabilities, and saving/loading model parameters.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Functions](#functions)
    - [Data Preprocessing](#data-preprocessing)
    - [Model Training](#model-training)
    - [Model Saving and Loading](#model-saving-and-loading)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Saurav-Navdhare/BayesianClassifier.git
    cd BayesianClassifier
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. **Data Preprocessing**: Convert raw data into a format suitable for training.
    ```go
    import "your-repo/utils"

    data := [][]string{
        {"Yes", "No", "Male"},
        {"No", "Yes", "Female"},
    }
    headers := []string{"Feature1", "Feature2", "Gender"}

    binaryData := utils.BinaryLabelling(data)
    df := utils.ConvertToDF(binaryData, headers)
    ```

2. **Train-Test Split**: Split the data into training and test sets.
    ```go
    import "BayesianClassifier/model"

    trainSet, testSet := model.TrainTestSplit(df, 0.2)
    ```

3. **Model Training**: Train the model using the training set.
    ```go
    classProbabilities, featureStats := model.CalculateProbabilities(trainSet)
    ```

4. **Save Model**: Save the trained model to a file.
    ```go
    model := model.Model{
        ClassProbabilities: classProbabilities,
        FeatureStats:       featureStats,
    }
    model.SaveModel("model.json", model)
    ```

5. **Load Model**: Load the model from a file and print the parameters.
    ```go
    loadedModel, err := model.LoadModel("model.json")
    if err != nil {
        log.Fatal(err)
    }
    ```

## Functions

### Data Preprocessing

- `BinaryLabelling`: Converts categorical data into binary labels.
- `ConvertToDF`: Converts a 2D slice of integers into a DataFrame-like map.

### Model Training

- `TrainTestSplit`: Splits the dataset into training and test sets.
- `CalculateProbabilities`: Calculates class probabilities and feature statistics.

### Model Saving and Loading

- `SaveModel`: Saves the model parameters to a JSON file.
- `LoadModel`: Loads the model parameters from a JSON file and prints them.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.