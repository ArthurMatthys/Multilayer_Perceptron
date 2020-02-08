package parsing

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

func ParseNucleus(filename string, train bool) ([]float64, []float64) {
	content := openFile(filename)
	lines := strings.Split(content, "\n")
	outputM := [2]float64{0, 1}
	outputB := [2]float64{1, 0}
	var output [][2]float64
	var allData [][30]float64
	for _, data := range lines {
		if data == "" {
			continue
		}
		line := strings.Split(data, ",")
		if line[1] == "M" {
			output = append(output, outputM)
		} else {
			output = append(output, outputB)
		}
		var inside [30]float64
		var err error
		for j, value := range line[2:] {
			inside[j], err = strconv.ParseFloat(value, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		allData = append(allData, inside)
	}
	if train {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(allData), func(i, j int) {
			allData[i], allData[j] = allData[j], allData[i]
			output[i], output[j] = output[j], output[i]
			lines[i], lines[j] = lines[j], lines[i]
		})
	}
	var retOutput []float64
	var retAllData []float64
	for _, value := range output {
		retOutput = append(retOutput, value[:]...)
	}
	for _, value := range allData {
		retAllData = append(retAllData, value[:]...)
	}
	return retAllData, retOutput
}
