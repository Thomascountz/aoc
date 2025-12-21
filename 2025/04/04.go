package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "04.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var totalRemovedRolls int

	grid := strings.Fields(string(data))
	removedRolls, grid := removeRolls(grid)
	for {
		if removedRolls == 0 {
			break
		}
		totalRemovedRolls += removedRolls
		removedRolls, grid = removeRolls(grid)
	}

	fmt.Println(totalRemovedRolls)
}

func removeRolls(grid []string) (int, []string) {
	var removed int
	var newGrid = make([]string, 0, len(grid))

	for rowIdx, row := range grid {
		var newRow strings.Builder

		for colIdx, char := range row {
			var neighbors int

			switch char {
			case '@':
				for i := -1; i <= 1; i++ {
					if rowIdx+i < 0 || rowIdx+i > len(grid)-1 {
						continue
					}
					for j := -1; j <= 1; j++ {
						if i == 0 && j == 0 {
							continue
						}

						if colIdx+j < 0 || colIdx+j > len(row)-1 {
							continue
						}

						if grid[rowIdx+i][colIdx+j] == '@' {
							neighbors++
						}
					}
				}

				if neighbors < 4 {
					removed++
					newRow.WriteRune('.')
				} else {
					newRow.WriteRune('@')
				}
			case '.':
				newRow.WriteRune('.')
			default:
				panic("unreachable")
			}
		}
		newGrid = append(newGrid, newRow.String())
	}
	return removed, newGrid
}
