package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInputString() (s string, err error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return s, err
	}

	s = string(bytes)

	return s, nil
}

func SumMetadataForNode(pos int, nums []int, sum *int) int {
	numChildren := nums[pos]
	numEntries := nums[pos+1]

	nextChildOffset := pos + 2
	for i := 0; i < numChildren; i++ {
		// visit each child of the node with depth first strategy
		nextChildOffset = SumMetadataForNode(nextChildOffset, nums, sum)
	}

	for j := 0; j < numEntries; j++ {
		*sum += nums[nextChildOffset+j]
	}

	return nextChildOffset + numEntries
}

func SumMetadata(nums []int) int {
	sum := 0

	SumMetadataForNode(0, nums, &sum)

	return sum
}

func FindRootNodeValue(nums []int) int {
	value, _ := FindNodeValue(0, nums)
	return value
}

func FindNodeValue(pos int, nums []int) (int, int) {
	numChildren := nums[pos]
	numEntries := nums[pos+1]
	nodeValue := 0

	nextChildOffset := pos + 2
	childValues := make([]int, numChildren)
	for i := 0; i < numChildren; i++ {
		// visit each child of the node with depth first strategy
		var childValue int
		childValue, nextChildOffset = FindNodeValue(nextChildOffset, nums)
		childValues[i] = childValue
	}

	if numChildren == 0 {
		for j := 0; j < numEntries; j++ {
			nodeValue += nums[nextChildOffset+j]
		}
	} else {
		for j := 0; j < numEntries; j++ {
			entry := nums[nextChildOffset+j]
			if entry > 0 && entry <= numChildren {
				nodeValue += childValues[entry-1]
			}
		}
	}

	return nodeValue, nextChildOffset + numEntries
}

func SplitStringIntoNumbers(s string) (result []int) {
	tokens := strings.Split(s, " ")

	for _, token := range tokens {
		num, _ := strconv.Atoi(token)
		result = append(result, num)
	}

	return result
}

func main() {
	input, err := ReadInputString()

	if err != nil {
		log.Fatal(err)
	}

	numbers := SplitStringIntoNumbers(input)

	fmt.Println("The sum of all metadata entries is:", SumMetadata(numbers))
	fmt.Println("The value of root node is:", FindRootNodeValue(numbers))
}
