package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, temp := range os.Args[1:] {
		ctemp, err := strconv.ParseFloat(temp, 64)

		if err == nil {
			ctemp := tempconv.Celsius(ctemp)
			ftemp := tempconv.CToF(ctemp)
			fmt.Printf("%s = %s", ctemp, ftemp)
		}
	}
}