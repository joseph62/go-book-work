package tempconv

import (
	"fmt"
	"os"
	"strconv"
)

type Pounds float64
type Kilograms float64

type Miles float64
type Kilometers float64

func main() {
	for _, value := range os.Args[1:] {
		scalar, err := strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Errorf("Failed to parse %s as a floating point number", value)
		}

		mi, km := Miles(scalar), Kilometers(scalar)

		fmt.Printf("%s = %s\n%s = %s\n", mi, MiToKm(mi), km, KmToMi(km))

		lbs, kg := Pounds(scalar), Kilograms(scalar)

		fmt.Printf("%s = %s\n%s = %s\n", lbs, LbsToKg(lbs), kg, KgToLbs(kg))
	}
}

func (p Pounds) String() string {
	return fmt.Sprintf("%g lbs", p)
}

func (kg Kilograms) String() string {
	return fmt.Sprintf("%g kg", kg)
}

func (m Miles) String() string {
	return fmt.Sprintf("%g mi", m)
}

func (km Kilometers) String() string {
	return fmt.Sprintf("%g km", km)
}

func KgToLbs(kg Kilograms) Pounds {
	return Pounds(kg * 2.2)
}

func LbsToKg(lbs Pounds) Kilograms {
	return Kilograms(lbs / 2.2)
}

func MiToKm(mi Miles) Kilometers {
	return Kilometers(mi * 1.609344)
}

func KmToMi(km Kilometers) Miles {
	return Miles(km / 1.609344)
}
