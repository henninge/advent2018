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

	fmt.Printf("Result: %s\n", findSimilar(serials))
}

func compare(first, second string) int {
	diffpos := -1
	for pos, firstChar := range first {
		if firstChar != rune(second[pos]) {
			if diffpos != -1 {
				return -1
			}
			diffpos = pos
		}
	}
	return diffpos
}

func findSimilar(serials []string) string {

	for first := 0; first < len(serials)-2; first++ {
		fmt.Printf("%d: %s\n", first, serials[first])
		for second := first + 1; second < len(serials)-1; second++ {
			fmt.Printf("  %d: %s\n", second, serials[second])
			pos := compare(serials[first], serials[second])
			if pos != -1 {
				result := serials[first]
				return fmt.Sprintf("%s%s", result[:pos], result[pos+1:])
			}
		}
	}
	return ""
}

func readFile() (value []string, error error) {
	file, err := ioutil.ReadFile("serials.txt")
	if error = err; err != nil {
		return
	}

	return strings.Split(string(file), "\n"), nil
}
