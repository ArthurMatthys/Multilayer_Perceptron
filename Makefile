all: build

dir:
	mkdir -p build/

build: dir
	go build -o build/multilayerPerceptron cmd/Multilayer_perceptron/main.go

fclean:
	rm -rf build

re: fclean all
