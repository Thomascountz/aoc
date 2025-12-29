package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "06.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	type Result struct {
		Addition       int
		Multiplication int
	}
	var results []Result
	var total int

	for scanner.Scan() {
		opIdx := 0
		line := scanner.Text()
		for operand := range strings.FieldsSeq(line) {
			switch operand {
			case "+":
				total += results[opIdx].Addition
			case "*":
				total += results[opIdx].Multiplication
			default:
				operandInt, _ := strconv.Atoi(operand)
				if opIdx > len(results)-1 {
					results = append(results, Result{operandInt, operandInt})
				} else {
					results[opIdx].Addition += operandInt
					results[opIdx].Multiplication *= operandInt
				}
			}
			opIdx++
		}
	}
	fmt.Println(total)
}
