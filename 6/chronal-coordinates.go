package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadCoordinatesFromFile() ([][]int, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var coords [][]int

	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ", ")

		x, _ := strconv.Atoi(tmp[0])
		y, _ := strconv.Atoi(tmp[1])

		xy := []int{x, y}
		coords = append(coords, xy)
	}

	return coords, scanner.Err()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func TotalDistanceAllCoordinates(coordinates [][]int, x int, y int) (distance int) {
	for _, coords := range coordinates {
		distance += ManhattanDistance(x, y, coords[0], coords[1])
	}

	return distance
}

func FindClosestLocation(locations [][]int, x int, y int) int {
	closestLocation := -1
	closestDistance := -1
	secondClosestDistance := -1
	for index, coords := range locations {
		dst := ManhattanDistance(x, y, coords[0], coords[1])

		if closestDistance < 0 || dst <= closestDistance {
			secondClosestDistance = closestDistance
			closestDistance = dst
			closestLocation = index
		}
	}

	if secondClosestDistance == closestDistance {
		// if two or more locations are closest to this point
		// then we return -1
		return -1
	}

	return closestLocation
}

func MaxNonInfinite(a []int, infiniteLocations []int) int {
	max := 0
	for pos, n := range a {
		if n > max && Index(infiniteLocations, pos) < 0 {
			max = n
		}
	}

	return max
}

func Index(haystack []int, needle int) int {
	for index, element := range haystack {
		if element == needle {
			return index
		}
	}

	return -1
}

func BottomRightPoint(coords [][]int) (x int, y int) {
	for _, point := range coords {
		if point[0] > x {
			x = point[0]
		}
		if point[1] > y {
			y = point[1]
		}
	}

	return x, y
}

func main() {
	locations, err := ReadCoordinatesFromFile()

	infiniteLocations := []int{}

	if err != nil {
		log.Fatal(err)
	}

	areas := make([]int, len(locations))
	right, bottom := BottomRightPoint(locations)

	safeArea := 0

	for i := 0; i < bottom; i++ {
		for j := 0; j < right; j++ {
			closestLocation := FindClosestLocation(locations, j, i)
			if closestLocation >= 0 {
				areas[closestLocation]++

				if (i == 0 || j == 0 || i == bottom-1 || j == right-1) && Index(infiniteLocations, closestLocation) < 0 {
					// we are on the border so the retrieved location
					// has infinite area
					infiniteLocations = append(infiniteLocations, closestLocation)
				}
			}

			// for second part
			totalDistance := TotalDistanceAllCoordinates(locations, j, i)
			fmt.Println(totalDistance)
			if totalDistance < 10000 {
				safeArea++
			}
		}
	}

	fmt.Println("The maximum non-infinite area is:", MaxNonInfinite(areas, infiniteLocations))
	fmt.Println("The safe area is:", safeArea)
}
