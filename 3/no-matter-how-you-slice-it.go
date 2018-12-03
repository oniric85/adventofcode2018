package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 1000x1000 fabric square
	var fabric [1000][1000]int
	total := 0

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("#[0-9]+ @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindStringSubmatch(line)

		// probably there's a more concise way to do this in Go..
		left, _ := strconv.Atoi(matches[1])
		top, _ := strconv.Atoi(matches[2])
		width, _ := strconv.Atoi(matches[3])
		height, _ := strconv.Atoi(matches[4])

		// update fabric
		for i := left; i < left+width; i++ {
			for j := top; j < top+height; j++ {
				fabric[i][j]++
				if fabric[i][j] == 2 {
					total++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of square inchese between two or more claims:", total)
}
