package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "05.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parsingRanges := true
	var freshIngredientIds []int
	var freshIngredientRanges [][2]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()

		if len(data) != 0 {
			if parsingRanges {
				before, after, found := strings.Cut(data, "-")
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

				freshIngredientRanges = append(freshIngredientRanges, [2]int{start, stop})
			} else { // parsingRanges == false
				id, err := strconv.Atoi(data)
				if err != nil {
					panic("parse fail (id)")
				}

				for _, freshIngredientRange := range freshIngredientRanges {
					if id >= freshIngredientRange[0] && id <= freshIngredientRange[1] {
						// fmt.Printf("%d is between %d and %d\n", id, freshIngredientRange[0], freshIngredientRange[1])
						freshIngredientIds = append(freshIngredientIds, id)
						break
					}
				}

			}
		} else {
			parsingRanges = false
		}
	}

	fmt.Println(len(freshIngredientIds))
}
