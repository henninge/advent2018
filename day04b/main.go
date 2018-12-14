package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type guardInfo struct {
	id            int
	minutesAsleep [60]int
}

func main() {
	logs, timestamps, err := readFile()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	guardStats := prcocessLogs(logs, timestamps)
	bestID, bestMinute := findBestGuard(guardStats)
	fmt.Printf("Result: %d", bestID*bestMinute)
}

var re = regexp.MustCompile("Guard #(\\d+) begins shift")

func prcocessLogs(logs map[string]string, timestamps []string) map[int]*guardInfo {

	guardStats := make(map[int]*guardInfo)

	var currentGuard *guardInfo
	var sleepMinute int
	var ok bool

	for _, timestamp := range timestamps {
		logLine := logs[timestamp]
		matches := re.FindStringSubmatch(logLine)

		fmt.Printf("%s %s\n", timestamp, logLine)
		if len(matches) != 0 {
			// Guard starts shift
			guardID, _ := strconv.Atoi(matches[1])
			currentGuard, ok = guardStats[guardID]
			if !ok {
				// Create new entry for this guard
				currentGuard = &guardInfo{id: guardID}
				guardStats[guardID] = currentGuard
			}
		} else {
			minute, _ := strconv.Atoi(timestamp[15:17])
			if logLine == "falls asleep" {
				// Remember when guard falls asleep
				sleepMinute = minute
			} else {
				// Guard wakes up, count the minutes.
				for m := sleepMinute; m < minute; m++ {
					currentGuard.minutesAsleep[m]++
				}
			}
		}
	}
	return guardStats
}

func findBestGuard(guardStats map[int]*guardInfo) (int, int) {
	var bestID, highestCount, bestMinute int
	for guardID, guard := range guardStats {
		minute, count := findBestMinute(guard.minutesAsleep[:])
		if count > highestCount {
			highestCount = count
			bestMinute = minute
			bestID = guardID
		}
	}
	fmt.Printf("Best minute(count): %d(%d)\n", bestMinute, highestCount)
	return bestID, bestMinute
}

func findBestMinute(minutes []int) (int, int) {
	var highestCount, bestMinute int
	for minute, count := range minutes {
		if count > highestCount {
			highestCount = count
			bestMinute = minute
		}
	}
	return bestMinute, highestCount
}

func readFile() (lines map[string]string, timestamps []string, error error) {
	file, err := os.Open("shifts.txt")
	if error = err; err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	lines = make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 3)
		if len(parts) != 3 {
			continue
		}
		timestamp := fmt.Sprintf("%s %s", parts[0], parts[1])
		lines[timestamp] = parts[2]
		timestamps = append(timestamps, timestamp)
	}
	error = scanner.Err()

	sort.Strings(timestamps)
	return
}
