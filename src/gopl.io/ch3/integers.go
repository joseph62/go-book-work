package main

import "fmt"

func main() {
	{
		var u uint8 = 255
		fmt.Println(u, u+1, u*u)

		var i int8 = 127
		fmt.Println(i, i+1, i*i)
	}

	{
		var x uint8 = 1<<1 | 1<<5
		var y uint8 = 1<<1 | 1<<2

		fmt.Printf("x       = %08b\n", x)
		fmt.Printf("y       = %08b\n", y)
		fmt.Printf("x  &  y = %08b\n", x&y)
		fmt.Printf("x  |  y = %08b\n", x|y)
		fmt.Printf("x  ^  y = %08b\n", x^y)
		fmt.Printf("x  &^ y = %08b\n", x&^y)

		for i := uint(0); i < 8; i++ {
			if x&(1<<i) != 0 {
				fmt.Printf("Position %v is populated in %08b\n", i, x)
			}
		}

		fmt.Printf("x  << 1 = %08b\n", x<<1)
		fmt.Printf("x  >> 1 = %08b\n", x>>1)
	}

	{
		// Operations like len() return ints which avoid
		// buffer underflows when for example iterating backwards through a slice

		medals := []string{"gold", "silver", "bronze"}

		for i := len(medals) - 1; i >= 0; i-- {
			fmt.Println(medals[i])
		}

		// unsigned integer types should be used for things like bit operations, hashing, and cryptography
	}

	{
		var apples int32 = 1
		var oranges int16 = 2
		// var compote int = apples + oranges << compilation error
		var compote int = int(apples) + int(oranges) // accepted

		fmt.Printf("Apples(%v) and Oranges(%v) combine to make %v\n", apples, oranges, compote)
	}

	{
		var f float64 = 3.141
		i := int(f)
		fmt.Printf("int(%v) = %v\n", f, i)
		f = 1.99
		fmt.Printf("int(%v) = %v\n", f, int(f))
	}

	{
		f := 1e100
		i := int(f)
		fmt.Printf("int(%v) = %v (Result depends on implementation)\n", f, i)
	}

	{
		o := 0666
		fmt.Printf("%d %[1]o %#[1]o\n", o)
		x := int64(0xdeadbeef)
		fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	}

	{
		ascii := 'a'
		unicode := rune(22269)
		newline := '\n'
		fmt.Printf("%d %[1]c %[1]q\n", ascii)
		fmt.Printf("%d %[1]c %[1]q\n", unicode)
		fmt.Printf("%d %[1]q\n", newline)
	}
}
