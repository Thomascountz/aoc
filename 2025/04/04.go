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

	var accessibleRolls int
	grid := strings.Fields(string(data))

	for rowNum, row := range grid {
		for colNum, char := range row {
			var neighbors int

			switch char {
			case '@':
				for i := -1; i <= 1; i++ {
					if rowNum+i < 0 || rowNum+i > len(row)-1 {
						continue
					}
					for j := -1; j <= 1; j++ {
						if i == 0 && j == 0 {
							continue
						}

						if colNum+j < 0 || colNum+j > len(grid)-1 {
							continue
						}

						if grid[rowNum+i][colNum+j] == '@' {
							neighbors++
						}
					}

				}

				if neighbors < 4 {
					// fmt.Printf("(%d, %d)\n", rowNum, colNum)
					accessibleRolls++
				}
			case '.':
				continue
			default:
				panic("unreachable")
			}
		}
	}

	fmt.Println(accessibleRolls)

}
