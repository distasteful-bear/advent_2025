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
func main() {
	wdir, err := os.Getwd()
	check(err)

	path := filepath.Join(wdir, "input.txt")
	data, err := os.ReadFile(path) // input.txt is < 1kb so you can read all into memory
	check(err)

	contents := string(data)
	lines := strings.Split(contents, "\n")
	fmt.Println("Hello World", lines)
}
