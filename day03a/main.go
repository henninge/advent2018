package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type claimInfo struct {
	id, lpos, tpos, width, height int
}

var fabric [1000][1000]int

func main() {
	claims, err := readFile()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	fillFabric(claims)
	//printFabric()
	fmt.Printf("Result: %d\n", countOverlaps())
}

func fillFabric(claims []claimInfo) {
	for _, claim := range claims {
		for tpos := claim.tpos; tpos < claim.tpos+claim.height; tpos++ {
			for lpos := claim.lpos; lpos < claim.lpos+claim.width; lpos++ {
				fabric[lpos][tpos]++
			}
		}
	}
}

func countOverlaps() (overlap int) {
	for _, fline := range fabric {
		for _, square := range fline {
			if square > 1 {
				overlap++
			}
		}
	}
	return
}

func printFabric() {
	for _, fline := range fabric {
		for _, square := range fline {
			if square == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", square)
			}
		}
		fmt.Print("\n")
	}
}

var re = regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")

func makeClaimInfo(line string) (claimInfo, error) {
	matches := re.FindStringSubmatch(line)

	claim := claimInfo{}
	if len(matches) == 0 {
		return claim, fmt.Errorf("Invalid line format: %s", line)
	}

	// Can safely ignore error because regex only collects digits.
	claim.id, _ = strconv.Atoi(matches[1])
	claim.lpos, _ = strconv.Atoi(matches[2])
	claim.tpos, _ = strconv.Atoi(matches[3])
	claim.width, _ = strconv.Atoi(matches[4])
	claim.height, _ = strconv.Atoi(matches[5])

	return claim, nil
}

func readFile() (value []claimInfo, error error) {
	file, err := ioutil.ReadFile("squares.txt")
	if error = err; err != nil {
		return
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		claim, err := makeClaimInfo(line)
		if error = err; err != nil {
			return
		}
		value = append(value, claim)
	}
	return
}
