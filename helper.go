package main

// contains all the information that defines a trained neural network
type neuralNet struct {
	config neuralNetConfig
	wHidden *mat.Dense
	bHidden *mat.Dense
	wOut	*mat.Dense
	bOut	*mat.Dense
}

// defines the neural network architecture and learning parameters
type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	numEpochs	  int
	learningRate  float64	
}