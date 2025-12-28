package main

import (
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(file)

	parsingRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRanges = false
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
		} else { // parsingRanges == false
			id, err := strconv.Atoi(line)
			if err != nil {
				panic("parse fail (id)")
			}

			for _, freshIngredientRange := range freshIngredientRanges {
				if id >= freshIngredientRange.Start && id <= freshIngredientRange.Stop {
					freshIngredientIdCount++
					break
				}
			}
		}
	}

	fmt.Printf("Fresh Ingredient Count: %d\n", freshIngredientIdCount)
}
