package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func replaceRuneInStr(s string, replace rune, index int) string {
	out := []rune(s)
	out[index] = replace
	return string(out)
}

func checkIsBale(b bale) bool {
	if 0 <= b.idx && b.idx < len(b.row) {
		return b.row[b.idx] == '@'
	}
	return false
}

type bale struct {
	row string
	idx int
}

func countBales(prevRow string, curRow string, nextRow string) (int, string) {
	availableBales := 0

	if len(prevRow) != len(curRow) || len(nextRow) != len(curRow) {
		panic(fmt.Errorf("Rows are not the same length!! PrevRow: %d CurRow: %d NextRow %d", len(prevRow), len(curRow), len(nextRow)))
	}

	var newRow = curRow

	for i, b := range curRow {

		hasBale := b == '@'
		if !hasBale {
			continue
		}

		totalAdjacentBales := 0

		adjacentBales := []bale{
			{row: prevRow, idx: i + 1},
			{row: prevRow, idx: i},
			{row: prevRow, idx: i - 1},

			{row: curRow, idx: i + 1},
			{row: curRow, idx: i - 1},

			{row: nextRow, idx: i + 1},
			{row: nextRow, idx: i},
			{row: nextRow, idx: i - 1},
		}

		for _, v := range adjacentBales {
			isBale := checkIsBale(v)
			if isBale {
				totalAdjacentBales += 1
			}
		}
		if totalAdjacentBales < 4 {
			availableBales += 1
			newRow = replaceRuneInStr(newRow, '.', i)
		}
	}
	return availableBales, newRow
}

func countThenRemoveBales(allBaleRows []string) (int, []string) {
	totalAvailableBales := 0
	var newBaleRows = []string{}

	blankRow := strings.Repeat(".", len(allBaleRows[0]))

	i := 0
	for {
		if i >= len(allBaleRows) {
			break
		}
		fmt.Printf("Row: %d Row Len: %d\n", i, len(allBaleRows))
		if len(allBaleRows[i]) == 0 {
			panic(fmt.Errorf("Found empty row! at Row Index: %d\n", i))
		}

		var prevRow string
		curRow := allBaleRows[i]
		var nextRow string

		if i == 0 {
			prevRow = blankRow
		} else {
			prevRow = allBaleRows[i-1]
		}
		if i == (len(allBaleRows) - 1) {
			nextRow = blankRow
		} else {
			nextRow = allBaleRows[i+1]
		}
		balesReadyForForkLift, newCurRow := countBales(prevRow, curRow, nextRow)
		totalAvailableBales += balesReadyForForkLift
		newBaleRows = append(newBaleRows, newCurRow)
		i += 1
	}
	fmt.Printf("Found bales in total: %d\n", totalAvailableBales)
	return totalAvailableBales, newBaleRows
}

func main() {

	wdir, err := os.Getwd()
	check(err)

	path := filepath.Join(wdir, "input.txt")

	fileContents, err := os.ReadFile(path)

	strData := string(fileContents)

	var rows = strings.Split(strData, "\n")
	if len(rows[len(rows)-1]) == 0 {
		rows = rows[:len(rows)-1] // last line is often blank
	}

	totalRemovedOnAllPasses := 0
	c := 0
	for {
		fmt.Printf("Pass Counter: %d\n", c)

		removedBales, rowsAfterPass := countThenRemoveBales(rows)
		rows = rowsAfterPass
		if removedBales > 0 {
			totalRemovedOnAllPasses += removedBales
		} else {
			break
		}
	}
	fmt.Printf("Total number of Passes: %d with a total %d total bales removed.\n", c, totalRemovedOnAllPasses)
}
