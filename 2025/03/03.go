package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "03.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const activeBatteryCount int = 12
	var totalJoltage int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batteryBank := scanner.Text()

		if len(batteryBank) == 0 {
			continue
		}

		totalJoltageForBank := 0
		searchStartIdx := 0
		for i := range activeBatteryCount {
			searchStopIdx := len(batteryBank) - (activeBatteryCount - 1 - i)

			joltageMaxIdx, joltageMax := findMaxJoltage(batteryBank, searchStartIdx, searchStopIdx)

			searchStartIdx = joltageMaxIdx + 1
			totalJoltageForBank = (totalJoltageForBank * 10) + joltageMax
		}

		totalJoltage += totalJoltageForBank

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(totalJoltage)
}

func findMaxJoltage(batteryBank string, searchStartIdx, searchStopIdx int) (int, int) {
	var joltageMaxIdx, joltageMax int
	for i := searchStartIdx; i < searchStopIdx; i++ {
		joltage := int(batteryBank[i] - '0')
		if joltage > joltageMax {
			joltageMax = joltage
			joltageMaxIdx = i
		}
	}
	return joltageMaxIdx, joltageMax
}
