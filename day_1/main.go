package main

import (
	"os"
	"path/filepath"
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

	path := filepath.Join(wdir, "input.txt") // input is 17kb so you can read into memory.
	data, err := os.ReadFile(path)
	check(err)

	contents := string(data)
	values := strings.Split(contents, "\n")

	for i, v := range values {

	}

}
