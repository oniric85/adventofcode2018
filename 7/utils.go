package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
)

func Shift(lst []string) (string, []string) {
	x, lst := lst[0], lst[1:]

	return x, lst
}

func Remove(s string, list []string) []string {
	for i := 0; i < len(list); i++ {
		if list[i] == s {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func FindReadySteps(instructions map[string][]string) (ready []string) {
	for step, before := range instructions {
		// a step is ready if the length of the associated array of steps is zero
		// this means that no step is needed before it can be carried on
		if len(before) == 0 {
			ready = append(ready, step)
		}
	}

	// return the list sorted to cope with puzzle constraints
	sort.Strings(ready)

	return ready
}

func ReadInstructionsFromFile() (instr map[string][]string, err error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return instr, err
	}

	defer file.Close()

	instr = make(map[string][]string)
	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
	for scanner.Scan() {
		matches := r.FindStringSubmatch(scanner.Text())

		if instr[matches[2]] == nil {
			instr[matches[2]] = []string{}
		}

		if instr[matches[1]] == nil {
			instr[matches[1]] = []string{}
		}

		instr[matches[2]] = append(instr[matches[2]], matches[1])
	}

	return instr, scanner.Err()
}
