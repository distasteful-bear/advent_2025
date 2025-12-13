package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	wdir, err := os.Getwd()
	check(err)

	path := filepath.Join(wdir, "input.txt")

	fileContents, err := os.ReadFile(path)

	strData := string(fileContents)

	allBanks := strings.Split(strData, "\n")

	maxJoltagesForAllBanks := 0

	for i, bank := range allBanks {
		if len(bank) == 0 {
			if i != (len(allBanks) - 1) {
				panic(fmt.Errorf("found an empty line which is not the final line of file (index: %d)", i))
			}
			continue
		}
		var maxJoltages = []rune{'0', '0'}
		var batteriesToUse = []int{0, 0}

		// find highest value
		for k, batteryJoltage := range bank {
			if batteryJoltage > maxJoltages[0] {
				maxJoltages[0] = batteryJoltage
				batteriesToUse[0] = k
			}
		}

		highestValueIsLastValue := batteriesToUse[0] == (len(bank) - 1)

		for k, batteryJoltage := range bank {
			if !highestValueIsLastValue && k <= batteriesToUse[0] {
				// you can skip all preceding values if the highest number found is not the last number.
				continue
			} else if k == batteriesToUse[0] {
				// if highest value is the end, the next highest number is always the correct value.
				continue
			}
			if batteryJoltage > maxJoltages[1] {
				maxJoltages[1] = batteryJoltage
				batteriesToUse[1] = k
			}
		}

		var joltageTotal string
		if batteriesToUse[0] < batteriesToUse[1] {
			joltageTotal = string(maxJoltages)
		} else {
			joltageTotal = string([]rune{maxJoltages[1], maxJoltages[0]})
		}

		maxJoltageForBank, err := strconv.Atoi(joltageTotal)
		check(err)

		maxJoltagesForAllBanks += maxJoltageForBank
		fmt.Printf("Max Joltage for Bank %d: %d with Batteries %d and %d\n\n", i, maxJoltageForBank, batteriesToUse[0], batteriesToUse[1])
	}
	fmt.Printf("Max Joltage for all Banks together: %d\n\n\n", maxJoltagesForAllBanks)
}
