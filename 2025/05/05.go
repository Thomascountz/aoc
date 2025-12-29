package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type IngredientRange struct {
	Start int
	End   int
}

type IngredientRangeList []IngredientRange

func (r IngredientRangeList) CollapseRanges() IngredientRangeList {
	// Sort
	sorted := slices.Clone(r)
	slices.SortFunc(sorted, func(a, b IngredientRange) int {
		return cmp.Compare(a.Start, b.Start)
	})

	// Merge
	ranges := make(IngredientRangeList, 0, len(sorted))
	ranges = append(ranges, sorted[0])

	for _, ingredientRange := range sorted[1:] {
		last := &ranges[len(ranges)-1]
		if ingredientRange.Start > last.End {
			ranges = append(ranges, ingredientRange)
		} else {
			last.End = max(ingredientRange.End, last.End)
		}
	}
	return ranges
}

func (r IngredientRangeList) CountIds() int {
	var ingredientIdsCount int
	for _, ingredientRange := range r {
		ingredientIdsCount += ingredientRange.End - ingredientRange.Start + 1
	}
	return ingredientIdsCount
}

func (r IngredientRangeList) Contains(id int) bool {
	_, found := slices.BinarySearchFunc(r, id, func(ir IngredientRange, target int) int {
		if ir.Start > target {
			return 1
		}
		if ir.End < target {
			return -1
		}
		return 0
	})
	return found
}

func main() {
	path := "05.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ingredientRangesRaw IngredientRangeList
	var ingredientRanges IngredientRangeList
	var ingredientCount int

	scanner := bufio.NewScanner(file)

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			ingredientRanges = ingredientRangesRaw.CollapseRanges()
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parsedRange := parseRange(line)
			ingredientRangesRaw = append(ingredientRangesRaw, parsedRange)
			continue
		} else {
			id := parseId(line)
			if ingredientRanges.Contains(id) {
				ingredientCount++
			}
		}
	}

	fmt.Printf("Fresh Ingredient Count: %d\n", ingredientCount)

	ingredientIdsCount := ingredientRanges.CountIds()
	fmt.Printf("Theoretical Fresh Ingredient Count: %d\n", ingredientIdsCount)
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
