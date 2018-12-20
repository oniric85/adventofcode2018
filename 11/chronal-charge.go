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

func SquarePower(top, left int, grid [300][300]int) (p int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			p += grid[top+i][left+j]
		}
	}

	return p
}

func CreateGrid(serialNumber int) (grid [300][300]int) {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			grid[i][j] = CellPower(j, i, serialNumber)
		}
	}

	return grid
}

func MaxSquare(grid [300][300]int) (maxX, maxY int) {
	maxPower := 0
	for i := 0; i < len(grid)-3; i++ {
		for j := 0; j < len(grid)-3; j++ {
			power := SquarePower(i, j, grid)

			if power > maxPower {
				maxPower = power
				maxY = i
				maxX = j
			}
		}
	}

	return maxX + 1, maxY + 1
}

func main() {
	grid := CreateGrid(6548)

	x, y := MaxSquare(grid)

	fmt.Println("Largest square:", x, ",", y)
}
