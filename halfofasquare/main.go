package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		for j := 5 - i; j > 0; j-- {
			fmt.Print("#")
		}
		fmt.Printf("\n")
		i++
	}
}
