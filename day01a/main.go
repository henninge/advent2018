package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	calibrations, err := readFile()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", calibrate(calibrations))
}

func calibrate(calibrations []int) int {
	sum := 0
	for _, calibration := range calibrations {
		sum = sum + calibration
	}
	return sum
}

func readFile() (value []int, error error) {
	file, err := ioutil.ReadFile("./resource/frequence.txt")
	if error = err; err != nil {
		return
	}

	for _, v := range strings.Split(string(file), "\n") {
		if v == "" {
			continue
		}

		elem, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		value = append(value, elem)
	}
	return
}
