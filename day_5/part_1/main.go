package main

import (
	"fmt"
	"log"
	"math/rand"
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

func quickSort(list []int64) []int64 {
	if len(list) <= 1 {
		return list
	}

	pivot := rand.Intn(len(list))
	pv := list[pivot]

	left := []int64{}
	right := []int64{}

	for i, v := range list {
		if i == pivot {
			continue
		}
		if v < pv {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return slices.Concat(left, []int64{pv}, right)
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
	idsToCheck := []int64{}

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
		} else {
			id, err := strconv.ParseInt(l, 10, 64)
			check(err)
			idsToCheck = append(idsToCheck, id)
		}
	}
	sortedIds := quickSort(idsToCheck)

	freshIds := []int64{}
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("Starting the double nested for loop.")
	for _, idRange := range idRanges {
		for _, id := range sortedIds {
			if slices.Contains(freshIds, id) {
				continue // skip if item status known
			}
			if id >= idRange.min && id <= idRange.max {
				freshIds = append(freshIds, id)
			} else {
				if id > idRange.max {
					break // range has no more ids worth checking bc asc sort
					// i checked this, and this if statement made it run in
					// .028 sec vs .042 sec
				}
			}
		}
	}
	log.Println("Finished the double nested for loop.")
	fmt.Printf("Found %d Fresh Items\n", len(freshIds))
}
