package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, temp := range os.Args[1:] {
		value, err := strconv.ParseFloat(temp, 64)

		if err == nil {
			ctemp := tempconv.Celsius(value)
			ftemp := tempconv.Fahrenheit(value)
			fmt.Printf("%s = %s\n", ctemp, tempconv.CToF(ctemp))
			fmt.Printf("%s = %s\n", ftemp, tempconv.FToC(ftemp))
		}
	}
}
