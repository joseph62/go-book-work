package popcount

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	for _, value := range os.Args[1:] {
		integer, err := strconv.Atoi(value)

		if err != nil {
			_ = fmt.Errorf("Failed to parse value %s into an integer: %v\n", value, err)
		} else {
			popcount := ExpressionPopCount(uint64(integer))
			fmt.Printf("Population of %d: %d\n", integer, popcount)
		}
	}
}

func LoopPopCount(x uint64) int {
	var population int

	for i := 0; i < 8; i++ {
		population += int(pc[byte(x>>(i*8))])
	}

	return population
}

func ExpressionPopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}