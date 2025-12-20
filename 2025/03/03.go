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

	totalJoltage := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		batteryBank := scanner.Text()
		firstIndex, firstJoltage := maxJoltage(batteryBank[:len(batteryBank)-1])
		_, secondJoltage := maxJoltage(batteryBank[firstIndex+1:])

		totalJoltageForBank := (firstJoltage * 10) + secondJoltage

		totalJoltage += totalJoltageForBank

	}
	fmt.Println(totalJoltage)
}

func maxJoltage(batteryBank string) (int, int) {
	var joltageMaxIndex, joltageMax int
	for i, battery := range batteryBank {
		joltage := int(battery - '0')
		if joltage > joltageMax {
			joltageMax = joltage
			joltageMaxIndex = i
		}
	}
	return joltageMaxIndex, joltageMax
}
