package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	s "strings"
)

func main() {
	dial := 50
	zeroCount := 0
	zeroPassCount := 0
	path := filepath.Join("01.txt")

	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Printf("%-4s\tdial:%-4d\tzero:%-4d\tzeroPass:%-4d\n", "X50", dial, zeroCount, zeroPassCount)
	for scanner.Scan() {
		instruction := scanner.Text()

		clicks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic("parse fail")
		}
		clicks = absInt(clicks)

		if s.HasPrefix(instruction, "L") {
			for click := range clicks {
				dial--
				if dial < 0 {
					dial = 99
					if click != 0 {
						zeroPassCount++
					}
				}
			}
		} else if s.HasPrefix(instruction, "R") {
			for click := range clicks {
				dial++
				if dial > 99 {
					dial = 0
					if click != clicks-1 {
						zeroPassCount++
					}
				}
			}
		} else {
			panic("unreachable")
		}

		if dial == 0 {
			zeroCount++
			zeroPassCount++
		}

		fmt.Printf("%-4s\tdial:%-4d\tzeros:%-4d\tzeroPasses:%-4d\n", instruction, dial, zeroCount, zeroPassCount)
	}
}

func absInt(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// Answers
// 01.a 1026
// 01.b 5923
