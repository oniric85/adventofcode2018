package main

import "fmt"

func CellPower(x, y, serialNumber int) (p int) {
	rackId := (x + 1) + 10
	p = rackId*(y+1) + serialNumber
	p *= rackId

	if p < 100 {
		p = 0
	} else {
		p = (p % 1000) / 100
	}

	p -= 5

	return p
}

func SquarePower(top, left int, grid [][]int, n int, powers [][]int) (p int) {
	for i := 0; i < n; i++ {
		if left > 0 {
			// use memoization to reduce time complexity
			p = powers[top][left-1]
			for k := 0; k < n; k++ {
				p -= grid[top+k][left-1]
				p += grid[top+k][left+n-1]
			}

			return p
		}
		for j := 0; j < n; j++ {
			p += grid[top+i][left+j]
		}
	}

	return p
}

func CreateGrid(serialNumber int, n int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			grid[i][j] = CellPower(j, i, serialNumber)
		}
	}

	return grid
}

func MaxSquare(grid [][]int, n int) (maxPower, maxX, maxY int) {
	powers := make([][]int, len(grid))

	for i := 0; i < len(grid)-n; i++ {
		powers[i] = make([]int, len(grid))
		for j := 0; j < len(grid)-n; j++ {
			power := SquarePower(i, j, grid, n, powers)
			powers[i][j] = power

			if power > maxPower {
				maxPower = power
				maxY = i
				maxX = j
			}
		}
	}

	return maxPower, maxX + 1, maxY + 1
}

func main() {
	grid := CreateGrid(6548, 300)

	maxPower := 0
	maxSize := 1
	maxX := 0
	maxY := 0
	for i := 1; i <= 300; i++ {
		power, x, y := MaxSquare(grid, i)
		if power > maxPower {
			maxPower = power
			maxSize = i
			maxX = x
			maxY = y
		}
	}

	fmt.Println("Largest square: size=", maxSize, ",power=", maxPower, ",x=", maxX, ",y=", maxY)
}
