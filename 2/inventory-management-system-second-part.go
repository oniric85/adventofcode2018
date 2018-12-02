package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func compareStrings(s string, t string) int {
	if len(s) != len(t) {
		return -1
	}

	diffs := 0
	posDiff := 0

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			diffs++
			posDiff = i
		}
	}

	if diffs == 1 {
		return posDiff
	}

	return -1
}

func main () {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	strings := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := scanner.Text()

		for _, str := range strings {
			if pos := compareStrings(str, x); pos >= 0 {
				fmt.Println("Found the two strings:", str, x)
				fmt.Println("They differ for character at position", pos)
				fmt.Println("The common part is:", str[:pos] + str[pos+1:])
				break
			}
		}

		strings = append(strings, x)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
}