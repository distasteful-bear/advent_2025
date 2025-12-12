package main

import (
	"fmt"
	"os"
)

func main() {
	i := 1
	for {
		if i > 25 {
			break
		}

		fmt.Printf("making dir %d\n", i)

		err := os.Mkdir(fmt.Sprintf("day_%d", i), 0777)

		if err != nil {
			fmt.Printf("\nerr: %v\n", err)
			break
		}

		i += 1
	}
}
