package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Let's do a double-liked list for fun.

type polymerUnit struct {
	uType          byte
	uPolarity      bool
	uName          string
	next, previous *polymerUnit
}

func (pu *polymerUnit) print() {
	var pol string
	if pu.uPolarity {
		pol = "+"
	} else {
		pol = "-"
	}
	fmt.Printf("%d%s", pu.uType, pol)
}

func (pu *polymerUnit) reactsWith(other *polymerUnit) bool {
	return pu.uType == other.uType && pu.uPolarity != other.uPolarity
}

func main() {
	polymerStart, err := readFile()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Length before: %d\n", countPolymer(polymerStart))
	polymerStart = reducePolymer(polymerStart)
	fmt.Printf("Length after: %d\n", countPolymer(polymerStart))
}

func countPolymer(polymerStart *polymerUnit) (count int) {
	polymerPos := polymerStart
	for polymerPos != nil {
		count++
		polymerPos = polymerPos.next
	}
	return
}

func printPolymer(polymerStart *polymerUnit) {
	polymerPos := polymerStart
	for polymerPos != nil {
		polymerPos.print()
		polymerPos = polymerPos.next
	}
}

func reducePolymer(polymerStart *polymerUnit) (polymer *polymerUnit) {
	polymer = polymerStart
	polymerPos := polymerStart
	//fmt.Printf("Result: %v", polymer)
	for polymerPos != nil {
		if polymerPos.next != nil && polymerPos.reactsWith(polymerPos.next) {
			if polymerPos.previous == nil {
				// Reaction happend at the beginning.
				// Remove the first two units and make the third the new start.
				polymerPos = polymerPos.next.next
				polymer = polymerPos
				if polymerPos != nil {
					polymerPos.previous = nil
				}
			} else {
				// Remove this unit and the next.
				// Move back to previous unit.
				polymerPos.previous.next = polymerPos.next.next
				polymerPos = polymerPos.previous
				if polymerPos.next != nil {
					// Update previous pointer if this is not the last unit.
					polymerPos.next.previous = polymerPos
				}
			}
		} else {
			// Advance to next unit
			polymerPos = polymerPos.next
		}
	}
	return
}

func readCode(rd io.Reader) (polymer *polymerUnit) {
	reader := bufio.NewReader(rd)

	var err error
	var uByte, uType byte
	var polarity bool
	var previousUnit *polymerUnit
	polymerPos := &polymer

	for {
		uByte, err = reader.ReadByte()
		if err != nil {
			break
		}

		// This is relying heavily on the ASCII order of things.
		if uByte < byte('A') || uByte > byte('z') {
			continue
		}

		if uByte > byte('Z') {
			uType = uByte - byte('a')
			polarity = true
		} else {
			uType = uByte - byte('A')
			polarity = false
		}
		*polymerPos = &polymerUnit{
			uType: uType, uName: string(uByte), uPolarity: polarity,
			previous: previousUnit}
		previousUnit = *polymerPos
		polymerPos = &previousUnit.next
	}
	return
}

func readFile() (polymer *polymerUnit, error error) {
	file, err := os.Open("polymer.txt")
	if error = err; err != nil {
		return
	}
	defer file.Close()

	return readCode(file), nil
}
