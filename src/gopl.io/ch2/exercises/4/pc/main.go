package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/exercises/popcount"
)

func main() {
	for _, value := range os.Args[1:] {
		integer, err := strconv.Atoi(value)

		if err != nil {
			_ = fmt.Errorf("Failed to parse value %s into an integer: %v\n", value, err)
		} else {
			popcount := popcount.ShiftingPopCount(uint64(integer))
			fmt.Printf("Population of %d: %d\n", integer, popcount)
		}
	}
}