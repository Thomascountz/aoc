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
	var dial int64 = 50
	zeroCount := 0

	path := filepath.Join("input", "01.txt")
	file, err := os.Open(path)

	if err != nil {
		panic("error reading file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("Instruction\tDial\tZeroCount")
	// fmt.Println("----------------------------------")

	for scanner.Scan() {
		instruction := scanner.Text()

		if s.HasPrefix(instruction, "L") {
			clicksInstruction := s.TrimPrefix(instruction, "L")
			clicks, err := strconv.ParseInt(clicksInstruction, 10, 64)
			if err != nil {
				panic("parse fail")
			}
			dial = mod(dial-clicks, 100)
			if dial == 0 {
				zeroCount++
			}
		} else if s.HasPrefix(instruction, "R") {
			clicksInstruction := s.TrimPrefix(instruction, "R")
			clicks, err := strconv.ParseInt(clicksInstruction, 10, 64)
			if err != nil {
				panic("parse fail")
			}
			dial = mod(dial+clicks, 100)
			if dial == 0 {
				zeroCount++
			}
		} else {
			panic("unreachable")
		}

		// fmt.Printf("%s\t\t%d\t\t%d\n", instruction, dial, zeroCount)

	}

	fmt.Println("----------------------------------")
	fmt.Printf("\t\t%d\t\t%d\n", dial, zeroCount)
}

func mod(a, b int64) int64 {
	return ((a % b) + b) % b
}

// Answers
// 01.a 1026
