package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readAndSortInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	return lines, scanner.Err()
}

func findSleepyGuard(m map[int][]int) int {
	sleepy := 0
	maxSleptMinutes := 0
	for guardId, minutes := range m {
		if len(minutes) > maxSleptMinutes {
			sleepy = guardId
			maxSleptMinutes = len(minutes)
		}
	}

	return sleepy
}

func findFrequentlySleptMinute(m map[int][]int) (int, int) {
	mostSleptMinute := 0
	maxTimesSlept := 0
	frequentSleeper := 0
	for guardId, minutes := range m {
		guardMostSleptMinute, timesSlept := findMostSleptMinute(minutes)
		if timesSlept > maxTimesSlept {
			mostSleptMinute = guardMostSleptMinute
			maxTimesSlept = timesSlept
			frequentSleeper = guardId
		}
	}
	return frequentSleeper, mostSleptMinute
}

func findMostSleptMinute(minutes []int) (int, int) {
	mostSleptMinute := 0
	maxSlept := 0

	minutesMap := make(map[int]int)

	for _, m := range minutes {
		minutesMap[m]++
		if minutesMap[m] > maxSlept {
			maxSlept = minutesMap[m]
			mostSleptMinute = m
		}
	}

	return mostSleptMinute, maxSlept
}

func main() {
	lines, err := readAndSortInput()

	guardsMap := make(map[int][]int)

	r, _ := regexp.Compile("[[0-9]{4}-([0-9]{2})-([0-9]{2}) ([0-9]{2}):([0-9]{2})] (.+)")
	s, _ := regexp.Compile("#([0-9]+)")

	currentGuard := 0
	startsSleeping := 0
	for _, line := range lines {
		lineMatches := r.FindStringSubmatch(line)
		guardMatches := s.FindStringSubmatch(lineMatches[5])

		minute, _ := strconv.Atoi(lineMatches[4])
		if guardMatches != nil {
			// found a guard ID
			currentGuard, _ = strconv.Atoi(guardMatches[1])
		}

		// toggle sleeping status
		if strings.Contains(lineMatches[5], "falls") {
			startsSleeping = minute
		} else if strings.Contains(lineMatches[5], "wakes") {
			// update sleeping minutes
			currentMinute, _ := strconv.Atoi(lineMatches[4])
			for i := startsSleeping; i < currentMinute; i++ {
				guardsMap[currentGuard] = append(guardsMap[currentGuard], i)
			}
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	// find the guard that sleeps the most
	sleepyGuard := findSleepyGuard(guardsMap)
	mostSleptMinute, _ := findMostSleptMinute(guardsMap[sleepyGuard])

	// first part solution
	fmt.Println("Guard with more slept minutes:", sleepyGuard)
	fmt.Println("The most slept minute of this guard is:", mostSleptMinute)
	fmt.Println("Product of the two values is:", sleepyGuard*mostSleptMinute)

	// second part solution
	frequentSleeperGuard, frequentlySleptMinute := findFrequentlySleptMinute(guardsMap)

	fmt.Println("Guard that slept the same minute the most:", frequentSleeperGuard)
	fmt.Println("The minute the guard spent sleeping the most:", frequentlySleptMinute)
	fmt.Println("Product of the two values is:", frequentSleeperGuard*frequentlySleptMinute)
}
