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

	var ingredientRangesRaw []IngredientRange
	var ingredientRanges []IngredientRange
	var ingredientCount int

	scanner := bufio.NewScanner(file)

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			ingredientRanges = mergeRanges(ingredientRangesRaw)

			ingredientCountMax := calculateMax(ingredientRanges)
			fmt.Printf("Theoretical Fresh Ingredient Count: %d\n", ingredientCountMax)
			assertEq("max fresh ingredients", ingredientCountMax, 354149806372909)

			parsingRanges = false
			continue
		}

		if parsingRanges {
			parsedRange := parseRange(line)
			ingredientRangesRaw = append(ingredientRangesRaw, parsedRange)
			continue
		} else {
			id := parseId(line)
			if idWithinRange(id, ingredientRanges) {
				ingredientCount++
			}
		}
	}

	fmt.Printf("Fresh Ingredient Count: %d\n", ingredientCount)
	assertEq("fresh ingredients", ingredientCount, 567)
	printMemUsage()
}

func mergeRanges(rangesRaw []IngredientRange) []IngredientRange {
	// Sort
	slices.SortFunc(rangesRaw, func(a, b IngredientRange) int {
		return cmp.Compare(a.Start, b.Start)
	})

	// Merge
	ranges := make([]IngredientRange, 0, len(rangesRaw))
	ranges = append(ranges, rangesRaw[0])

	for _, ingredientRange := range rangesRaw[1:] {
		last := &ranges[len(ranges)-1]
		if ingredientRange.Start > last.Stop {
			ranges = append(ranges, ingredientRange)
		} else {
			last.Stop = max(ingredientRange.Stop, last.Stop)
		}
	}
	return ranges
}

func calculateMax(ranges []IngredientRange) int {
	var ingredientCountMax int
	for _, ingredientRange := range ranges {
		ingredientCountMax += ingredientRange.Stop - ingredientRange.Start + 1
	}
	return ingredientCountMax
}

func parseRange(line string) IngredientRange {
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

	return IngredientRange{start, stop}
}

func parseId(line string) int {
	id, err := strconv.Atoi(line)
	if err != nil {
		panic("parse fail (id)")
	}
	return id
}

func idWithinRange(id int, ranges []IngredientRange) bool {
	return slices.ContainsFunc(ranges, func(ingredientRange IngredientRange) bool {
		return id >= ingredientRange.Start && id <= ingredientRange.Stop
	})
}

func assertEq(desc string, got int, want int) {
	if want != got {
		fmt.Printf("FAIL - %s\n\tgot: %-15d\twant: %-15d\tdiff: %-15d\n", desc, got, want, got-want)
	} else {
		fmt.Printf("PASS - %s\n\tgot: %-15d\n", desc, got)
	}
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
