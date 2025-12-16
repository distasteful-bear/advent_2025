package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// int64 needed for these values
// (which int should work for on fedora, but use it explicitly for cross platform)
type idSet struct {
	max int64
	min int64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func mergeOverlappingSets(idRanges []idSet) []idSet {

	mergedRanges := []idSet{}
	indexesAlreadyMerged := []int{}

	for i, r := range idRanges {
		if slices.Contains(indexesAlreadyMerged, i) {
			continue
		}

		mergedRanges = append(mergedRanges, r)
		r1 := r

		c := 0
		for {
			if c >= len(idRanges) {
				break
			}
			if c == i || slices.Contains(indexesAlreadyMerged, c) {
				c += 1
				continue
			}
			r2 := idRanges[c]
			if (r2.min >= r1.min && r2.min <= r1.max) || (r2.max >= r1.min && r2.max <= r1.max) {
				r1 = idSet{
					max: max(r1.max, r2.max),
					min: min(r1.min, r2.min),
				}
				mergedRanges[len(mergedRanges)-1] = r1
				indexesAlreadyMerged = append(indexesAlreadyMerged, c)
				c = 0
			}
			c += 1
		}
	}

	return mergedRanges
}

func main() {
	wdir, err := os.Getwd()
	check(err)

	path := filepath.Join(wdir, "self_test.txt")
	data, err := os.ReadFile(path) // input.txt is < 1kb so you can read all into memory
	check(err)

	contents := string(data)
	lines := strings.Split(contents, "\n")

	idRanges := []idSet{}

	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		if strings.Contains(l, "-") {
			ids := strings.Split(l, "-")
			if len(ids) != 2 {
				panic(fmt.Errorf("expected line %d to have two ids separated by '-'", i))
			}
			minId, err := strconv.ParseInt(ids[0], 10, 64)
			check(err)

			maxId, err := strconv.ParseInt(ids[1], 10, 64)
			check(err)

			idRanges = append(idRanges, idSet{
				min: minId,
				max: maxId,
			})
		}
	}

	idRangesWithoutOverlap := mergeOverlappingSets(idRanges)
	totalUniqueValuesInAllRanges := 0
	for _, r := range idRangesWithoutOverlap {
		fmt.Println(r.min, r.max)
		valuesInRange := (r.max - r.min) + 1 // count is inclusive (fence posts not boards)
		totalUniqueValuesInAllRanges += int(valuesInRange)
	}

	fmt.Printf("Found %d items considered fresh by id ranges\n", totalUniqueValuesInAllRanges)
}
