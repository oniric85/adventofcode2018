package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Index(vs []int, n int) int {
	for i, v := range vs {
		if v == n {
			return i
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 1000x1000 fabric square
	var fabric [1000][1000]int

	intactClaims := []int{}

	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindStringSubmatch(line)

		// probably there's a more concise way to do this in Go..
		id, _ := strconv.Atoi(matches[1])
		left, _ := strconv.Atoi(matches[2])
		top, _ := strconv.Atoi(matches[3])
		width, _ := strconv.Atoi(matches[4])
		height, _ := strconv.Atoi(matches[5])

		isIntact := true

		// update fabric
		for i := left; i < left+width; i++ {
			for j := top; j < top+height; j++ {
				if fabric[i][j] > 0 {
					// the current claim is certainly not intact
					isIntact = false
					// not the first visit of this square, mark the claim as not intact
					index := Index(intactClaims, fabric[i][j])
					if index >= 0 {
						intactClaims = append(intactClaims[:index], intactClaims[index+1:]...)
					}
				}
				// first visit of this square, mark it with the id of the current claim
				fabric[i][j] = id
			}
		}

		if isIntact {
			intactClaims = append(intactClaims, id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The only intact claim is:", intactClaims[0])
}
