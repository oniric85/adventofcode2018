package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	freq := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		freq += x
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	fmt.Println("Final frequency:", freq)
}