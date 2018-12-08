package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func buildFreqSlice() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	freqSlice := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		freqSlice = append(freqSlice, x)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return freqSlice
}

func findFirstRepeatedFreq() int {
	freqSlice := buildFreqSlice()
	freqMap := make(map[int]bool)

	// initialize map
	freq := 0
	freqMap[freq] = true

	for {
		for i := 0; i < len(freqSlice); i++ {
			freq += freqSlice[i]

			if _, ok := freqMap[freq]; ok {
				return freq
			}

			freqMap[freq] = true
		}
	}
}

func main() {
	fmt.Println("First repetition:", findFirstRepeatedFreq())
}
