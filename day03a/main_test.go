package main

import (
	"testing"
)

type regexTest struct {
	line string
}

func TestRegex(t *testing.T) {
	line := "#1124 @ 613,170: 17x"
	_, err := makeClaimInfo(line)
	if err == nil {
		t.Errorf("Should Fail: %s\n", line)
	}

	info, _ := makeClaimInfo("#1124 @ 613,170: 17x28\n")
	expected := claimInfo{id: 1124, lpos: 613, tpos: 170, width: 17, height: 28}

	if info != expected {
		t.Errorf("%+v\n", info)
	}

	info, _ = makeClaimInfo("#1 @ 3,0: 3x5\n")
	expected = claimInfo{id: 1, lpos: 3, tpos: 0, width: 3, height: 5}

	if info != expected {
		t.Errorf("%+v\n", info)
	}
}
