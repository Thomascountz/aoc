package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type IngredientRange struct {
	Start int
	Stop  int
}

func main() {
	path := "05.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var freshIngredientRanges []IngredientRange
	var freshIngredientIdCount int
	var freshIngredientIdCountMax int

	scanner := bufio.NewScanner(file)

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fmt.Printf("Before Merge: %d\n", len(freshIngredientRanges))

			// Sort
			slices.SortFunc(freshIngredientRanges, func(a, b IngredientRange) int {
				return cmp.Compare(a.Start, b.Start)
			})

			// Merge
			for i := 1; i < len(freshIngredientRanges); i++ {
				// If the current starts at or before the previous stops
				if freshIngredientRanges[i].Start <= freshIngredientRanges[i-1].Stop {
					// If the current stops at or after the previous stops
					if freshIngredientRanges[i].Stop >= freshIngredientRanges[i-1].Stop {
						// Extend the previous' stop to that of the current
						freshIngredientRanges[i-1].Stop = freshIngredientRanges[i].Stop
					}
					// Remove the current
					freshIngredientRanges = append(freshIngredientRanges[:i], freshIngredientRanges[i+1:]...)
					// Take a step back
					i -= 1
				}
			}

			// Calc Theoretical Max
			for _, freshIngredientRange := range freshIngredientRanges {
				freshIngredientIdCountMax += (freshIngredientRange.Stop - freshIngredientRange.Start) + 1
			}

			parsingRanges = false
			fmt.Printf("After Merge: %d\n", len(freshIngredientRanges))
			continue
		}

		if parsingRanges {
			before, after, found := strings.Cut(line, "-")
			if !found {
				panic("parse fail (Cut)")
			}

			start, err := strconv.Atoi(before)
			if err != nil {
				panic("parse fail (start)")
			}

			stop, err := strconv.Atoi(after)
			if err != nil {
				panic("parse fail (stop)")
			}

			freshIngredientRanges = append(freshIngredientRanges, IngredientRange{start, stop})
		} else { // !parsingRanges

			id, err := strconv.Atoi(line)
			if err != nil {
				panic("parse fail (id)")
			}

			if slices.ContainsFunc(freshIngredientRanges, func(freshIngredientRange IngredientRange) bool {
				return id >= freshIngredientRange.Start && id <= freshIngredientRange.Stop
			}) {
				freshIngredientIdCount++
			}
		}
	}

	fmt.Printf("Fresh Ingredient Count: %d\n", freshIngredientIdCount)
	fmt.Printf("Theoretical Fresh Ingredient Count: %d\n", freshIngredientIdCountMax)
	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("Total Mallocs: %d\n", m.Mallocs)
}
