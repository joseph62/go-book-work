package main

import (
	"fmt"
	"math"
)

func main() {
	{
		var f float64 = 16777216
		fmt.Println(f == f+1)
	}

	const (
		e        = 2.71828
		Avogadro = 6.02214129e23
		Planck   = 6.62606957e-34
	)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d  eˣ = %8.3f\n", x, math.Exp(float64(x)))
	}

	{
		var z float64
		fmt.Println(z, -z, 1/z, -1/z, z/z)
	}

	{
		nan := math.NaN()
		fmt.Println(nan == nan, nan < nan, nan > nan)
	}

	{
		compute := func(n float64) (float64, bool) {
			result := n / n

			if math.IsNaN(result) || math.IsInf(result, int(n)) {
				return 0, false
			}

			return result, true
		}

		value, ok := compute(0.0)

		fmt.Print("Zero divded by Zero results in a ")

		if ok {
			fmt.Println("good value:", value)
		} else {
			fmt.Println("bad value:", value)
		}
	}
}
