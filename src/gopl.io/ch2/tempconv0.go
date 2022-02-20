package main

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("absolute zero = %s or %s\n", CToF(AbsoluteZeroC), AbsoluteZeroC)
	fmt.Printf("freezing point = %s or %s\n", CToF(FreezingC), FreezingC)
	fmt.Printf("boiling point = %s or %s\n", CToF(BoilingC), BoilingC)

	fmt.Println(AbsoluteZeroC.String())
	fmt.Printf("%v\n", AbsoluteZeroC)
	fmt.Printf("%s\n", AbsoluteZeroC)
	fmt.Println(AbsoluteZeroC)
	fmt.Printf("%g\n", AbsoluteZeroC)
	fmt.Println(float64(AbsoluteZeroC))
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
