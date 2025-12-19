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
		rangeBeginEnd := s.Split(r, "-")

		rangeBegin, err := strconv.Atoi(s.TrimSuffix(rangeBeginEnd[0], "\n"))
		check(err)
		rangeEnd, err := strconv.Atoi(s.TrimSuffix(rangeBeginEnd[1], "\n"))
		check(err)

		for i := rangeBegin; i <= rangeEnd; i++ {
			id := strconv.Itoa(i)

			if isInvalid(id) {
				invalidSum += i
			}
		}
	}
	fmt.Println(invalidSum)
}

func isInvalid(id string) bool {
	idLength := len(id)
	for stride := 1; stride < idLength; stride++ {
		if idLength%stride != 0 {
			continue
		}

		var segments []string
		for i := 0; i < idLength; i += stride {
			segment := id[i : i+stride]
			segments = append(segments, segment)
		}

		invalid := true
		for i := 1; i < len(segments); i++ {
			if segments[i] != segments[i-1] {
				invalid = false
				break
			}
		}

		if invalid {
			return invalid
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
