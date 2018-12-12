package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	serials, err := readFile()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", checkSum(serials))
}

func checkSum(serials []string) int {

	sum2 := 0
	sum3 := 0

	for _, serial := range serials {
		charCount := map[rune]int{}
		for _, char := range serial {
			charCount[char]++
		}

		found2 := false
		found3 := false
		for _, count := range charCount {
			if count == 2 {
				found2 = true
			}
			if count == 3 {
				found3 = true
			}
		}
		if found2 {
			sum2++
		}
		if found3 {
			sum3++
		}
	}
	return sum2 * sum3
}

func readFile() (value []string, error error) {
	file, err := ioutil.ReadFile("serials.txt")
	if error = err; err != nil {
		return
	}

	return strings.Split(string(file), "\n"), nil
}
