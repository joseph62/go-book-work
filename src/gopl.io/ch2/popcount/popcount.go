package main

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
			fmt.Errorf("Failed to parse value %s into an integer: %v\n", value, err)
		} else {
			popcount := PopCount(uint64(integer))
			fmt.Printf("Population of %d: %d\n", integer, popcount)
		}
	}
}

/*
	000 	pc[0] + 0x00000000 & 0x00000001 = 0
	001 	pc[0] + 0x00000001 & 0x00000001 = 1
	002 	pc[1] + 0x00000010 & 0x00000001 = 1
	003 	pc[1] + 0x00000011 & 0x00000001 = 2
	004 	pc[2] + 0x00000100 & 0x00000001 = 1
	005 	pc[2] + 0x00000101 & 0x00000001 = 2
	006 	pc[3] + 0x00000110 & 0x00000001 = 2
	007 	pc[3] + 0x00000111 & 0x00000001 = 3
	008 	pc[4] + 0x00001000 & 0x00000001 = 1
	009 	pc[4] + 0x00001001 & 0x00000001 = 2
	010 	pc[5] + 0x00001010 & 0x00000001 = 2
	011 	pc[5] + 0x00001011 & 0x00000001 = 3
	012 	pc[6] + 0x00001100 & 0x00000001 = 2
*/

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
