package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func readInputString() ([]byte, error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func checkUnits(first byte, second byte) bool {
	if unicode.ToLower(rune(first)) == unicode.ToLower(rune(second)) {
		if first != second {
			return true
		}

		return false
	}

	return false
}

func findRemainingUnits(polymer []byte) string {
	i := 0
	for i < len(polymer)-1 {
		if checkUnits(polymer[i], polymer[i+1]) {
			polymer = append(polymer[:i], polymer[i+2:]...)
			// backtrack a single step to verify if a new reaction was created
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}

	return string(polymer)
}

func main() {
	units, err := readInputString()

	if err != nil {
		log.Fatal(err)
	}

	remainingUnits := findRemainingUnits(units)

	fmt.Println("The length of the remaining units is:", len(remainingUnits))
}
