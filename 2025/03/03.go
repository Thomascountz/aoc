package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	path := "03.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const activatedBatteries int = 12
	var totalJoltage int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batteryBank := scanner.Text()

		totalJoltageForBank := 0
		searchStartIndex := 0
		for i := range activatedBatteries {
			searchStopIndex := (len(batteryBank)) - (activatedBatteries - 1 - i)
			// fmt.Printf("battery:%s\tstart:2%d\tstop:2%d\n", batteryBank, searchStartIndex, searchStopIndex)

			joltageMaxBatteryIndex, joltageMax := findMaxJoltage(batteryBank, searchStartIndex, searchStopIndex)

			searchStartIndex = joltageMaxBatteryIndex + 1
			totalJoltageForBank += joltageMax * int((math.Pow10(activatedBatteries - 1 - i)))
		}

		totalJoltage += totalJoltageForBank

	}
	fmt.Println(totalJoltage)
}

func findMaxJoltage(batteryBank string, searchStartIndex, searchStopIndex int) (int, int) {
	var joltageMaxIndex, joltageMax int
	for i := searchStartIndex; i < searchStopIndex; i++ {
		battery := batteryBank[i]
		joltage := int(battery - '0')
		if joltage > joltageMax {
			joltageMax = joltage
			joltageMaxIndex = i
		}
	}
	return joltageMaxIndex, joltageMax
}
