package parsing

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func openFile(filename string) string {
	file, _ := os.Open(filename)
	stat, _ := os.Stat(filename)

	b := make([]byte, stat.Size())
	_, err := file.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	var content string
	content = string(b)
	return content
}

//func ParseNucleus(filename string) (input, output []float64) {
func ParseNucleus(filename string) {
	content := openFile(filename)
	lines := strings.Split(content, "\n")
	outputM := [2]uint{0, 1}
	outputB := [2]uint{1, 0}
	var output [][2]uint
	//	var allData [][30]float64
	for _, data := range lines {
		if data == "" {
			continue
		}
		line := strings.Split(data, ",")

		fmt.Println(line)
		if line[1] == "M" {
			output = append(output, outputM)
		} else {
			output = append(output, outputB)
		}
		//		inside := data[2:]
	}
	fmt.Println(output)
}
