package main

import (
	"testing"
)

type compareTest struct {
	first    string
	second   string
	expected int
}

func TestCompare(t *testing.T) {

	testPairs := []compareTest{
		compareTest{"abcd", "abcd", -1},
		compareTest{"abeecd", "abffcd", -1},
		compareTest{"abcde", "acbde", -1},
		compareTest{"abcde", "ebcda", -1},
		compareTest{"abcde", "bbcde", 0},
		compareTest{"abcde", "abcdf", 4},
		compareTest{"abcde", "abtde", 2},
	}
	for _, ct := range testPairs {
		pos := compare(ct.first, ct.second)
		if pos != ct.expected {
			t.Errorf("%s, %s, %d != %d ", ct.first, ct.second, pos, ct.expected)
		}
	}
}

type similarTest struct {
	serials  []string
	expected string
}

func TestSimilar(t *testing.T) {

	testPairs := []similarTest{
		similarTest{[]string{"abcdef", "abcddd", "abcdde"}, "abcdd"},
		similarTest{[]string{"abvdef", "abcdef", "adddff"}, "abdef"},
		similarTest{[]string{"abcdef", "zudhfd", "uuuiii"}, ""},
		similarTest{[]string{"ubcdef", "abcdef", "adddff"}, "bcdef"},
		similarTest{[]string{"ubidef", "abcdef", "adddff", "aacdef"}, "acdef"},
	}
	for _, st := range testPairs {
		result := findSimilar(st.serials)
		if result != st.expected {
			t.Errorf("%+v, %s != %s ", st.serials, result, st.expected)
		}
	}
}
