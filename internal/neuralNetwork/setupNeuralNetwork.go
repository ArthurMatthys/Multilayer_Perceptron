package neuralNetwork

import (
	"fmt"
	"github.com/ArthurMatthys/Multilayer_Perceptron/internal/parsing"
	"gonum.org/v1/gonum/mat"
)

type NeuralNetwork struct {
	Regularisation float64
	Alpha          float64
	Epoch          uint64
	Save           string
	Train          bool
	Weights        []mat.Dense
	Layer          []Layer
}

func structure() {
}

func Setup(mP NeuralNetwork, datafile, structure, load string) {
	output, data := parsing.ParseNucleus(datafile, mP.Train)
	denseOutput := mat.NewDense(len(output)/30, 30, output)
	denseData := mat.NewDense(len(data)/2, 2, data)
	fmt.Println(denseData, denseData.At(568, 0))
	fmt.Println(denseOutput, denseOutput.At(568, 29))
	//	if load != "" {
	//		importNetwork(mP, data, load)
	//	} else {
	//		createNetwork(mP, data, structure)
	//	}
}
