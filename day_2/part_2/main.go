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

type idRange struct {
	min int
	max int
}

func checkOnlyRepetitions(s string, l int) bool {
	substringsOfLen := []string{}
	c := 0
	for {
		substringsOfLen = append(substringsOfLen, s[c:c+l])
		c += l
		if c >= len(s) {
			break
		}
	}

	allValuesIdentical := true
	for _, v := range substringsOfLen {
		if v != substringsOfLen[0] {
			allValuesIdentical = false
		}
	}
	return allValuesIdentical
}

func isIdValid(id int) bool {
	s := strconv.Itoa(id)

	repeatTextLength := 1
	for {
		if repeatTextLength > (len(s) / 2) { // this is intentional integer divion
			break
		}
		if len(s)%repeatTextLength == 0 {
			onlyRepitions := checkOnlyRepetitions(s, repeatTextLength)
			if onlyRepitions {
				return false
			}
		}
		repeatTextLength += 1
	}
	return true
}

func main() {
	wdir, err := os.Getwd()
	check(err)

	path := filepath.Join(wdir, "input.txt")
	data, err := os.ReadFile(path) // input.txt is < 1kb so you can read all into memory
	check(err)

	contents := string(data)
	lines := strings.Split(contents, "\n")
	sets := strings.Split(lines[0], ",") // only the first line needs to be worked on

	var idRanges []idRange

	for i, v := range sets {
		setIdxs := strings.Split(v, "-")

		if len(setIdxs) != 2 {
			panic(fmt.Errorf("index set did not have 2 elements, set num: %d", i))
		}
		minIdx, minErr := strconv.Atoi(setIdxs[0])
		maxIdx, maxErr := strconv.Atoi(setIdxs[1])

		check(minErr)
		check(maxErr)

		idRanges = append(idRanges, idRange{
			min: minIdx,
			max: maxIdx,
		})
	}

	var totalOfInvalidIds int = 0
	for i, r := range idRanges {

		var invalidIds []int
		curId := r.min
		for {
			if curId > r.max {
				break
			}
			if !isIdValid(curId) {
				invalidIds = append(invalidIds, curId)
			}
			curId += 1
		}
		fmt.Printf("%d Invalid Ids found in range (%d): %d - %d\n", len(invalidIds), i, r.min, r.max)
		for _, id := range invalidIds {
			totalOfInvalidIds += id
		}
	}
	fmt.Printf("Total of all Invalid Ids: %d\n\n", totalOfInvalidIds)
}
