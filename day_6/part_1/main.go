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

type opToken struct {
	op     rune // * +
	values []int
}

func tokenizeValues(s string) []int {
	var values []int
	var buffer []rune
	for _, c := range s {
		if c != ' ' {
			buffer = append(buffer, c)
		}
		if c == ' ' && len(buffer) != 0 {
			s := string(buffer)
			newValue, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			values = append(values, newValue)
			buffer = []rune{}
		}
	}
	if len(buffer) > 0 {
		s := string(buffer)
		newValue, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		values = append(values, newValue)
	}
	return values
}

func tokenizeOperations(s string) []rune {
	var values []rune
	for _, r := range s {
		if r == ' ' {
			continue
		}
		values = append(values, r)
	}
	return values
}

func parseTokens(valueLines []string, operationLine string) []opToken {
	tokenValueLines := [][]int{}

	for _, l := range valueLines {
		tokens := tokenizeValues(l)
		tokenValueLines = append(tokenValueLines, tokens)
	}

	tokenOperators := tokenizeOperations(operationLine)

	if len(tokenOperators) != len(tokenValueLines[0]) {
		panic(fmt.Errorf("Operators and Lines are not the same length!\nToken Operators: %d\nTokenValues: %d\n", len(tokenOperators), len(tokenValueLines[0])))
	}

	numOpTokens := len(tokenOperators)

	var ops []opToken

	i := 0
	for {
		if i == numOpTokens {
			break
		}
		values := []int{}
		for _, l := range tokenValueLines {
			values = append(values, l[i])
		}
		ops = append(ops, opToken{
			op:     tokenOperators[i],
			values: values,
		})
		i += 1
	}
	return ops
}
func calculateResult(o opToken) int {
	if o.op == '*' {
		total := 1
		for _, v := range o.values {
			total *= v
		}
		return total
	} else if o.op == '+' {
		total := 0
		for _, v := range o.values {
			total += v
		}
		return total
	} else {
		panic(fmt.Errorf("Unknown Operator!! - %v", o.op))
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

	valueLines := []string{}
	operationLine := ""

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "*") {
			operationLine = line
		} else {
			valueLines = append(valueLines, line)
		}
	}

	operationTokens := parseTokens(valueLines, operationLine)
	grandTotal := 0

	for i, o := range operationTokens {
		t := calculateResult(o)
		grandTotal += t
		fmt.Printf("Total for index: %v = %v\n", i, t)
	}
	fmt.Printf("Grand Total: %v\n", grandTotal)
}
