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

func checkOverlap(set1 idSet, set2 idSet) bool {
	// this seems pretty verbose, but all of these are required
	// if set2 is a superset of set1 the last 2 if statements catch it.
	if set2.min >= set1.min && set2.min <= set1.max {
		return true
	}
	if set2.max >= set1.min && set2.max <= set1.max {
		return true
	}
	if set1.min >= set2.min && set1.min <= set2.max {
		return true
	}
	if set1.max >= set2.min && set1.max <= set2.max {
		return true
	}
	return false
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
			if checkOverlap(r1, r2) {
				r1 = idSet{
					min: min(r1.min, r2.min),
					max: max(r1.max, r2.max),
				}
				mergedRanges[len(mergedRanges)-1] = r1
				indexesAlreadyMerged = append(indexesAlreadyMerged, c)
				// the change to the current range might change the overlap status of previous id ranges
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

	path := filepath.Join(wdir, "input.txt")
	data, err := os.ReadFile(path) // input.txt is < 1kb so you can read all into memory
	check(err)

	contents := string(data)
	lines := strings.Split(contents, "\n")

	idRanges := []idSet{}

	for i, l := range lines {
		if len(l) == 0 {
			break
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

		fmt.Println("Range for Id Set (post Merge):", r.min, " - ", r.max)
		totalUniqueValuesInAllRanges += int(valuesInRange)
		fmt.Println("Num Ids in Range: ", valuesInRange)
	}

	fmt.Printf("Found %d items considered fresh by id ranges\n", totalUniqueValuesInAllRanges)
}
