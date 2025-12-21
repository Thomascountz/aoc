package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
)

func main() {
	path := "04.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var totalRemovedRolls int

	grid := bytes.Fields(data)
	rows := len(grid)
	cols := len(grid[0])

	gridCurr := make([][]byte, rows)
	gridNext := make([][]byte, rows)

	for i, line := range grid {
		gridCurr[i] = make([]byte, cols)
		copy(gridCurr[i], line)
		gridNext[i] = make([]byte, cols)
		copy(gridNext[i], line)
	}

	for {
		removedRolls, changed := removeRolls(gridCurr, gridNext)
		if !changed {
			break
		}

		totalRemovedRolls += removedRolls
		gridCurr, gridNext = gridNext, gridCurr
	}

	fmt.Println(totalRemovedRolls)
	printMemUsage()
}

func removeRolls(sourceGrid, destGrid [][]byte) (int, bool) {
	var removed int
	var changed bool
	for rowIdx, row := range sourceGrid {
		for colIdx, char := range row {
			var neighbors int

			switch char {
			case '@':
				for i := -1; i <= 1; i++ {
					if rowIdx+i < 0 || rowIdx+i > len(sourceGrid)-1 {
						continue
					}
					for j := -1; j <= 1; j++ {
						if i == 0 && j == 0 {
							continue
						}

						if colIdx+j < 0 || colIdx+j > len(row)-1 {
							continue
						}

						if sourceGrid[rowIdx+i][colIdx+j] == '@' {
							neighbors++
						}
					}
				}

				if neighbors < 4 {
					removed++
					changed = true
					destGrid[rowIdx][colIdx] = '.'
				} else {
					destGrid[rowIdx][colIdx] = '@'
				}
			case '.':
				destGrid[rowIdx][colIdx] = '.'
			default:
				panic("unreachable")
			}
		}
	}
	return removed, changed
}

// For info on each, see: https://golang.org/pkg/runtime/#MemStats
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("Total Mallocs: %d\n", m.Mallocs)
}
