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

func calcMaxJoltage(line string) (int, error) {

	maxJoltage := [12]rune{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}

	remainingBatteries := len(line) - 1
	lastFinishedBattery := 0

	for _, battery := range line {
		c := 0
		for {
			if c >= 12 {
				break
			}
			if lastFinishedBattery+1 == c {
				maxJoltage[c] = battery
				lastFinishedBattery = c
				break
			}
			if battery > maxJoltage[c] && remainingBatteries >= (11-c) {
				maxJoltage[c] = battery
				lastFinishedBattery = c
				break
			}
			c += 1
		}
		remainingBatteries -= 1
	}

	strJoltage := ""
	for _, r := range maxJoltage {
		strJoltage += string(r)
	}
	return strconv.Atoi(strJoltage)
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

		joltage, err := calcMaxJoltage(bank)
		check(err)

		fmt.Printf("Max Joltage found for line %d: %d\n", i, joltage)
		maxJoltagesForAllBanks += joltage
	}
	fmt.Printf("Max Joltage for all Banks together: %d\n\n\n", maxJoltagesForAllBanks)
}
