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
	path := filepath.Join("input", "01.txt")

	file, err := os.Open(path)
	if err != nil {
		panic("error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instruction := scanner.Text()

		clicks, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic("parse fail")
		}

		if s.HasPrefix(instruction, "L") {
			dial = mod(dial-clicks, 100)
		} else if s.HasPrefix(instruction, "R") {
			dial = mod(dial+clicks, 100)
		} else {
			panic("unreachable")
		}

		if dial == 0 {
			zeroCount++
		}

		fmt.Printf("%s\tdial:%4d\t%4d\n", instruction, dial, zeroCount)

	}

	fmt.Printf("\nDial: %d\tZeros: %d\n", dial, zeroCount)
}

func mod(a, b int) int {
	// return ((a % b) + b) % b
	return a % b
}

// Answers
// 01.a 1026
