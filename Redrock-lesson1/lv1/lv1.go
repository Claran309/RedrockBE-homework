package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%3d * %d = %-2d", j, i, i*j)
		}
		fmt.Print("\n")
	}
}
