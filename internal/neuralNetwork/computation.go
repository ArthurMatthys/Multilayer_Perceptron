package neuralNetwork

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

type Layer struct {
	size       int
	weigth     mat.Dense
	activation func(mat.Dense) mat.Dense
	derivate   func(mat.Dense) mat.Dense
}

func createWeight() {
}

func computeCost(ouputTarget, output mat.Dense) float64 {
	var cost float64
	denseLog := func(i, j int, value float64) float64 {
		return math.Log(value)
	}
	output.Apply(denseLog, &output)
	return cost
}

func forward() {
}

func backward() {
}

func train(trainSet, trainOut, testSet, testOut mat.Dense, network NeuralNetwork) {
}

func test(input, output mat.Dense, network NeuralNetwork) {
}
