package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadInputString() ([]byte, error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func CheckUnits(first byte, second byte) bool {
	if ToLower(first) == ToLower(second) {
		if first != second {
			return true
		}

		return false
	}

	return false
}

func ToLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}

	return b
}

func Index(haystack []byte, needle byte) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}

	return -1
}

func FindUniqueTypes(polymer []byte) []byte {
	result := []byte{}

	for _, t := range polymer {
		b := ToLower(t)
		if Index(result, b) == -1 {
			result = append(result, b)
		}
	}

	return result
}

func RemoveTypeFromPolymer(polymer []byte, t byte) []byte {
	result := []byte{}

	for _, c := range polymer {
		unit := ToLower(c)
		if unit != t {
			result = append(result, c)
		}
	}

	return result
}

func FindRemainingUnits(polymer []byte) string {
	str := make([]byte, len(polymer))
	copy(str, polymer)

	i := 0
	for i < len(str)-1 {
		if CheckUnits(str[i], str[i+1]) {
			str = append(str[:i], str[i+2:]...)
			// backtrack a single step to verify if a new reaction was created
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}

	return string(str)
}

func main() {
	units, err := ReadInputString()

	if err != nil {
		log.Fatal(err)
	}

	remainingUnits := FindRemainingUnits(units)

	fmt.Println("The length of the input polymer after all the reactions is:", len(remainingUnits))

	// now we remove a type at a time from the polymer and react it
	for _, t := range FindUniqueTypes(units) {
		fmt.Println("Checking after removing type", string(t))
		polymer := RemoveTypeFromPolymer(units, t)
		remainingUnits = FindRemainingUnits(polymer)

		fmt.Println("The length of the reduced polymer is:", len(remainingUnits))
	}
}
