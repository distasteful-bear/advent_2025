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

func isIdValid(id int) bool {
	s := strconv.Itoa(id)

	if len(s)%2 != 0 {
		return true
	}

	substr1 := s[:len(s)/2]
	substr2 := s[len(s)/2:]

	return substr1 != substr2
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
