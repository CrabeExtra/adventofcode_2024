package main

/*
*
	Is there a more efficient way to do this? Almost definitely. But I had a busy day. big O(width*height) for both answers
*/

import (
	"day4/helpers"
	"day4/lib"
	"fmt"
)

func partOne(data [][]string) {
	totalXmas := 0

	for rowIndex, row := range data {
		var stack lib.Stack

		for colIndex, letter := range row {
			// for each letter, check for X, if not X continue
			if letter == "X" {
				// moving clockwise from rowIndex-1 colIndex-1. through + + and then finally rowIndex, colIndex-1
				// if found, add to the stack
				for i := rowIndex - 1; i <= rowIndex+1; i++ {
					for j := colIndex - 1; j <= colIndex+1; j++ {
						if !(i == rowIndex && j == colIndex) && lib.InBounds(i, j, data) {

							// if the letter is M, add to the stack
							if data[i][j] == "M" {
								stack = stack.Push(map[string]int{"rowPos": i, "colPos": j})
							}
						}
					}
				}

				// now check the stack move in the direction of the difference between the row/col and the rowPos, colPos
				for i := len(stack) - 1; i >= 0; i-- {
					var entry map[string]int
					stack, entry = stack.Pop()

					rowDiff := entry["rowPos"] - rowIndex
					colDiff := entry["colPos"] - colIndex

					ARow := entry["rowPos"] + rowDiff
					ACol := entry["colPos"] + colDiff

					SRow := entry["rowPos"] + (rowDiff * 2)
					SCol := entry["colPos"] + (colDiff * 2)

					// if within array and rest of word exists?
					if lib.InBounds(ARow, ACol, data) && lib.InBounds(SRow, SCol, data) && data[ARow][ACol] == "A" && data[SRow][SCol] == "S" {

						totalXmas++
					}
				}
			}

		}
	}
	fmt.Println("Part one total:")
	fmt.Println(totalXmas)
}

func partTwo(data [][]string) {
	totalMAS := 0

	// Loop through the grid
	for rowIndex, row := range data {
		for colIndex, letter := range row {
			// Only start checking if we find an "M"
			// if the top left and bottom right are inbounds, then any adjacent letter is inbounds (otherwise there can't be a cross)
			if letter == "A" && lib.InBounds(rowIndex-1, colIndex-1, data) && lib.InBounds(rowIndex+1, colIndex+1, data) {

				topLeftRow := rowIndex - 1
				topLeftCol := colIndex - 1

				topRightRow := rowIndex - 1
				topRightCol := colIndex + 1

				bottomLeftRow := rowIndex + 1
				bottomLeftCol := colIndex - 1

				bottomRightRow := rowIndex + 1
				bottomRightCol := colIndex + 1

				// check back diagonal
				if lib.BackDiagonalMas(topLeftRow, topLeftCol, bottomRightRow, bottomRightCol, data) && lib.ForeDiagonalMas(topRightRow, topRightCol, bottomLeftRow, bottomLeftCol, data) {
					totalMAS++
				}

			}
		}
	}

	fmt.Println("Part two total:")
	fmt.Println(totalMAS)
}

func main() {
	data := helpers.ReadInput()

	partOne(data)
	partTwo(data)
}
