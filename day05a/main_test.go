package main

import (
	"strings"
	"testing"
)

func TestReadCodeEmpty(t *testing.T) {
	code := ""
	result := readCode(strings.NewReader(code))
	if result != nil {
		t.Errorf("Empty string returned: %+v.", *result)
	}
}

func compareUnits(unit *polymerUnit, uType byte, uPolarity bool) bool {
	return unit != nil && unit.uType == uType && unit.uPolarity == uPolarity
}

func TestReadCodeA(t *testing.T) {
	code := "A"
	result := readCode(strings.NewReader(code))
	if !compareUnits(result, 0, false) {
		t.Errorf("%s returned: %+v.", code, *result)
	}
}

func TestReadCodea(t *testing.T) {
	code := "a"
	result := readCode(strings.NewReader(code))
	if !compareUnits(result, 0, true) {
		t.Errorf("%s returned: %+v.", code, *result)
	}
}

func TestReadCodeZ(t *testing.T) {
	code := "Z"
	result := readCode(strings.NewReader(code))
	if !compareUnits(result, 25, false) {
		t.Errorf("%s returned: %+v.", code, *result)
	}
}

func TestReadCodez(t *testing.T) {
	code := "z"
	result := readCode(strings.NewReader(code))
	if !compareUnits(result, 25, true) {
		t.Errorf("%s returned: %+v.", code, *result)
	}
}

func TestReadCodes2(t *testing.T) {
	code := "AB"
	result := readCode(strings.NewReader(code))
	success := (compareUnits(result, 0, false) &&
		compareUnits(result.next, 1, false) &&
		result.next.previous == result)
	if !success {
		t.Errorf("%s returned: %+v, %+v.", code, *result, *result.next)
	}
}

func TestReadCodes3(t *testing.T) {
	code := "Aca"
	result := readCode(strings.NewReader(code))
	success := (compareUnits(result, 0, false) &&
		compareUnits(result.next, 2, true) &&
		compareUnits(result.next.next, 0, true) &&
		result.next.previous == result &&
		result.next.next.previous == result.next)
	if !success {
		t.Errorf("%s returned: %+v, %+v.", code, *result, *result.next)
	}
}

func TestReduction(t *testing.T) {
	// These code should all reduce to the same "Aca"
	codes := []string{"Aca", "bBAca", "BbAca", "AcbBa", "AbBca", "AcaBb", "BbAbdDBcBDdbaBb"}
	for _, code := range codes {
		result := reducePolymer(readCode(strings.NewReader(code)))
		success := (compareUnits(result, 0, false) &&
			compareUnits(result.next, 2, true) &&
			compareUnits(result.next.next, 0, true) &&
			result.next.previous == result &&
			result.next.next.previous == result.next)
		if !success {
			t.Errorf("%s returned: %+v, %+v.", code, *result, *result.next)
		}
	}
}
