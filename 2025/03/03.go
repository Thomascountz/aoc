package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "03.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalJoltage := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		batteries := strings.Split(bank, "")

		firstIndex, firstJoltage := maxJoltageIndex(batteries[:len(batteries)-1])
		_, secondJoltage := maxJoltageIndex(batteries[firstIndex+1:])

		totalJoltageForBank := (firstJoltage * 10) + secondJoltage

		totalJoltage += totalJoltageForBank

	}
	fmt.Println(totalJoltage)
}

func maxJoltageIndex(batteries []string) (index int, maxJoltage int) {
	index = 0
	maxJoltage = 0
	for i, battery := range batteries {
		joltage, err := strconv.Atoi(battery)
		if err != nil {
			panic(err)
		}
		if joltage > maxJoltage {
			maxJoltage = joltage
			index = i
		}
	}
	return
}
