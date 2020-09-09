package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filename string) (grid [9][9]string) {
	dat, _ := ioutil.ReadFile(filename)
	grid = makeGrid(dat)
	return
}

func makeGrid(fileData []byte) (grid [9][9]string) {
	rows := strings.Split(string(fileData), "\n")
	for rowInd, row := range rows {

		for coulInd, col := range strings.Split(row, "") {
			grid[rowInd][coulInd] = col
		}

	}
	return
}

func displayGrid(grid [9][9]string) {
	fmt.Println()
	rowCounter := 0

	for _, val := range grid {
		if rowCounter == 3 {
			fmt.Println("---------------")
			rowCounter = 0
		}
		inRowCounter := 0
		for _, value := range val {

			inRowCounter++

			switch inRowCounter {
			case 3:
				fmt.Print(value)
				fmt.Print(" | ")
			case 9:
				fmt.Print(value)
				fmt.Print("\n")
			case 6:
				fmt.Print(value)
				fmt.Print(" | ")
			default:
				fmt.Print(value)
			}

		}

		rowCounter++
	}
	fmt.Println()
}

func solve(grid [9][9]string) (solvedGrid [9][9]string) {

	for rowInd, row := range grid {
		for colInd, col := range row {
			if col == "." {
				variables := findPossibleValues(grid, [2]int{rowInd, colInd})
				if len(variables) == 1 {
					grid[rowInd][colInd] = variables[0]
				} else {
					continue
				}

			}
		}
	}

	if findEmptyPosition(grid) {
		grid = solve(grid)
	}

	return grid
}

func getRow(grid [9][9]string, position [2]int) (row [9]string) {
	row = grid[position[0]]
	return
}

func getColumn(grid [9][9]string, position [2]int) (column [9]string) {
	for rowInd, row := range grid {
		for _, col := range row[position[1]] {
			column[rowInd] = string(col)
		}
	}
	return
}

func getBlock(grid [9][9]string, position [2]int) (block [9]string) {

	var blockRows []int
	counter := 0

	switch position[0] {
	case 0, 1, 2:
		blockRows = []int{0, 1, 2}
	case 3, 4, 5:
		blockRows = []int{3, 4, 5}
	case 6, 7, 8:
		blockRows = []int{6, 7, 8}
	}

	switch position[1] {
	case 0, 1, 2:
		for _, val := range blockRows {
			for _, value := range grid[val][:3] {
				block[counter] = value
				counter++
			}
		}
	case 3, 4, 5:
		for _, val := range blockRows {
			for _, value := range grid[val][3:6] {
				block[counter] = value
				counter++
			}
		}
	case 6, 7, 8:
		for _, val := range blockRows {
			for _, value := range grid[val][6:9] {
				block[counter] = value
				counter++
			}
		}
	}

	return
}

func findPossibleValues(grid [9][9]string, position [2]int) (result []string) {
	r := getRow(grid, position)
	c := getColumn(grid, position)
	b := getBlock(grid, position)

	allowedNums := "123456789"
	have := ""

	for _, val := range r {
		have += val
	}

	for _, val := range c {
		have += val
	}

	for _, val := range b {
		have += val
	}

	for _, val := range allowedNums {
		if !strings.Contains(have, string(val)) {
			result = append(result, string(val))
		}
	}

	return
}

func findEmptyPosition(grid [9][9]string) bool {

	for _, row := range grid {
		for _, col := range row {
			if col == "." {
				return true
			}
		}
	}
	return false
}

// postion = (row,column)

func main() {

	grid := readFile("sudoku.txt")
	displayGrid(grid)
	solved := solve(grid)
	displayGrid(solved)
}
