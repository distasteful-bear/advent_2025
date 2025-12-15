package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(idx int, s string, rotation int) int {

	if len(s) < 2 {
		panic(fmt.Errorf("could not parse line %d because it has less than 2 characters", idx))
	}
	v, l := utf8.DecodeRuneInString(s)
	prefix := v
	var err error
	rotation, err = strconv.Atoi(s[l:])
	check(err)

	if prefix == 'L' {
		rotation = rotation * -1
	}

	fmt.Printf("Line Input: %s\n", s)
	fmt.Printf("Output Int: %d\n", rotation)
	return rotation
}

func main() {
	wdir, err := os.Getwd()
	check(err)

	path := filepath.Join(wdir, "input.txt")
	data, err := os.ReadFile(path) // input.txt is 17kb so you can read all into memory
	check(err)

	contents := string(data)
	values := strings.Split(contents, "\n")

	var zeroIndexes int = 0
	var rotation int = 0
	var curIndex int = 50

	numValues := len(values)
	for i, v := range values {
		if i+1 >= numValues && len(v) < 2 {
			// skip the last row which is a blank
			continue
		}

		rotation = parseLine(i, v, rotation)

		curIndex = curIndex + rotation

		for {
			if curIndex < 0 {
				curIndex += 100
			} else if curIndex > 99 {
				curIndex -= 100
			} else {
				break
			}
		}
		fmt.Printf("Current Index After Rotation: %d \n\n", curIndex)

		if curIndex == 0 {
			zeroIndexes += 1
		}
	}

	fmt.Printf("Number of Zero Indexes: %d\n", zeroIndexes)
}
