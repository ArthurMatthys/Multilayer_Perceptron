package setup

import ()

type NeuralNetwork struct {
	Regularisation float64
	Alpha          float64
	Epoch          uint64
	Save           string
	Train          bool
}

func Setup(mP NeuralNetwork, data, structure, load string) {

}
