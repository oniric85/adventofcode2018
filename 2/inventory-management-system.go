package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func reverseMap(m map[rune]int) map[int]bool {
    n := make(map[int]bool)
    for _, v := range m {
        n[v] = true
    }
    return n
}

func checkLetters(s string) (bool, bool) {
	lettersMap := make(map[rune]int)

	for _, char := range s {
		lettersMap[char] += + 1
	}

	reversedMap := reverseMap(lettersMap)

	_, has2 := reversedMap[2]
	_, has3 := reversedMap[3]

	return has2, has3
}

func main () {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count2, count3 := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := scanner.Text()
		res2, res3 := checkLetters(x)
		if res2 {
			count2++
		}
		if res3 {
			count3++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Checksum:", count2 * count3)
}