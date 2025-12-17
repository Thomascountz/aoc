package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	s "strings"
)

func main() {
	path := filepath.Join("02.txt")
	data, err := os.ReadFile(path)
	check(err)

	invalidSum := 0
	for r := range s.SplitSeq(string(data), ",") {
		rangeBeginrangeEnd := s.Split(r, "-")

		rangeBegin, err := strconv.Atoi(s.TrimSuffix(rangeBeginrangeEnd[0], "\n"))
		check(err)
		rangeEnd, err := strconv.Atoi(s.TrimSuffix(rangeBeginrangeEnd[1], "\n"))
		check(err)

		for i := rangeBegin; i <= rangeEnd; i++ {
			id := strconv.Itoa(i)

			idLength := len(id)
			if idLength%2 != 0 {
				continue // Skip if id cannnot be symetrical
			}

			// "slice" operator: [low(inclusive):high(exclusive)]
			left := id[:(idLength / 2)]
			right := id[(idLength / 2):]

			if left == right {
				invalidSum += i
			}
		}
	}
	fmt.Println(invalidSum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
