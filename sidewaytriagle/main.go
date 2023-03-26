package main

import "fmt"

func main() {
	for row := 4; row > 0; row-- {
		for hashNum := 1; hashNum <= 5-row; hashNum++ {
			fmt.Print("#")
		}

		fmt.Printf("\n")
	}
	for row := 1; row <= 3; row++ {
		for hashNum := 1; hashNum <= 4-row; hashNum++ {
			fmt.Print("#")
		}

		fmt.Printf("\n")
	}
}
