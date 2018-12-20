package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x  int
	y  int
	vx int
	vy int
}

func ReadPointsFromFile() ([]Point, error) {
	file, err := os.Open("input.txt")

	points := []Point{}

	if err != nil {
		return points, err
	}

	defer file.Close()

	r, _ := regexp.Compile("^position=< *(-?[0-9]+), *(-?[0-9]+)> velocity=< *(-?[0-9]+), *(-?[0-9]+)>$")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := r.FindStringSubmatch(scanner.Text())
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])

		p := Point{x: x, y: y, vx: vx, vy: vy}
		points = append(points, p)
	}

	return points, nil
}

func CalculateEdges(points []Point) (int, int, int, int) {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for i := 0; i < len(points); i++ {
		if points[i].x < minX {
			minX = points[i].x
		}
		if points[i].y < minY {
			minY = points[i].y
		}
		if points[i].x > maxX {
			maxX = points[i].x
		}
		if points[i].y > maxY {
			maxY = points[i].y
		}
	}

	return minX, maxX, minY, maxY
}

func NormalizePoints(points []Point, offsetX int, offsetY int) []Point {
	for i := 0; i < len(points); i++ {
		points[i].x -= offsetX
		points[i].y -= offsetY
	}

	return points
}

func PrintMessage(points []Point) {
	minX, maxX, minY, maxY := CalculateEdges(points)

	points = NormalizePoints(points, minX, minY)

	grid := make([][]bool, -minY+maxY+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, -minX+maxX+1)
	}

	for i := 0; i < len(points); i++ {
		grid[points[i].y][points[i].x] = true
	}

	// adjusting parameters to precisely show the message
	for i := 109; i < len(grid); i++ {
		for j := 190; j < len(grid[i]); j++ {
			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	points, err := ReadPointsFromFile()

	if err != nil {
		log.Fatal(err)
	}

	// iteration was found by looking around iterations that provided
	// the minimum perimeter of the bounding box of all points
	iteration := 10054
	for j := 0; j < len(points); j++ {
		points[j].x += points[j].vx * iteration
		points[j].y += points[j].vy * iteration
	}

	PrintMessage(points)
}
