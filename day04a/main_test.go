package main

import (
	"testing"
)

var testTimestamps = []string{
	"[2018-12-02 23:59]",
	"[2018-12-03 00:30]",
	"[2018-12-03 00:35]",
	"[2018-12-03 23:59]",
	"[2018-12-04 00:20]",
	"[2018-12-04 00:35]",
	"[2018-12-04 23:59]",
	"[2018-12-05 00:05]",
	"[2018-12-05 00:06]",
	"[2018-12-05 00:16]",
	"[2018-12-05 00:18]",
	"[2018-12-06 00:05]",
	"[2018-12-06 00:10]",
	"[2018-12-06 00:31]",
	"[2018-12-07 00:00]",
	"[2018-12-07 00:58]",
	"[2018-12-07 00:59]",
}

var testLogs = map[string]string{
	"[2018-12-02 23:59]": "Guard #44 begins shift",
	"[2018-12-03 00:30]": "falls asleep",
	"[2018-12-03 00:35]": "wakes up",
	"[2018-12-03 23:59]": "Guard #45 begins shift",
	"[2018-12-04 00:20]": "falls asleep",
	"[2018-12-04 00:35]": "wakes up",
	"[2018-12-04 23:59]": "Guard #46 begins shift",
	"[2018-12-05 00:05]": "falls asleep",
	"[2018-12-05 00:06]": "wakes up",
	"[2018-12-05 00:16]": "falls asleep",
	"[2018-12-05 00:18]": "wakes up",
	"[2018-12-06 00:05]": "Guard #44 begins shift",
	"[2018-12-06 00:10]": "falls asleep",
	"[2018-12-06 00:31]": "wakes up",
	"[2018-12-07 00:00]": "Guard #46 begins shift",
	"[2018-12-07 00:58]": "falls asleep",
	"[2018-12-07 00:59]": "wakes up",
}

func TestGuards(t *testing.T) {
	guard := prcocessLogs(testLogs, testTimestamps)
	if guard.id != 44 {
		t.Errorf("Wrong guard: %d\n", guard.id)
	}
	if guard.totalAsleep != 26 {
		t.Errorf("Wrong totalASleep: %+v\n", guard)
	}
	// TODO: Test minutesAsleep
}
