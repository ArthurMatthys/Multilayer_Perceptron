package main

import (
	"flag"
	"github.com/ArthurMatthys/Multilayer_Perceptron/internal/neuralNetwork"
	"log"
	"os"
)

func check_file(filename string) {
	file, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	if file.Mode().IsDir() {
		msg := "stat " + filename + ": is a directory"
		log.Fatalf(msg)
	}
}

func main() {
	var data, structure string
	var load string
	var mP neuralNetwork.NeuralNetwork

	flag.StringVar(&data, "d", "", "Filename of the data file")
	flag.StringVar(&structure, "s", "", "Structure of the")
	flag.StringVar(&mP.Save, "save", "", "Name of the file to save the weigths")
	flag.StringVar(&load, "load", "", "Name of the file to load the weigths")
	flag.Float64Var(&mP.Regularisation, "r", 0.016,
		"The value of the regularisation in the cost function")
	flag.Float64Var(&mP.Alpha, "a", 0, "The value of alpha.")
	flag.Uint64Var(&mP.Epoch, "e", 1000, "Number of epoch")
	flag.BoolVar(&mP.Train, "p", true, "If you want to predict or train")
	flag.Parse()

	if data == "" || structure == "" {
		log.Fatalf(`You need to precise the filename for both of the data file and the structure of Neural Network`)
	}
	if load != "" && structure != "" {
		log.Fatalf("If loading file is specify, you should give a structure")
	}
	if !mP.Train && load == "" {
		log.Fatalf("If you want to predict something, you need to load a trained Multilayer Perceptron")
	}
	check_file(data)
	check_file(structure)
	if mP.Save != "" {
		check_file(mP.Save)
	}
	if load != "" {
		check_file(load)
	}
	if mP.Regularisation < 0 {
		log.Fatalf("The regularisation cannot be negative")
	}
	if mP.Alpha < 0 {
		log.Fatalf("The alpha cannot be negative")
	}
	if mP.Epoch < 0 {
		log.Fatalf("The number of epoch cannot be negative")
	}
	neuralNetwork.Setup(mP, data, structure, load)
}
