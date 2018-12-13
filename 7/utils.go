package main

import (
	"bufio"
	"os"
	"regexp"
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
