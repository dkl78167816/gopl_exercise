package main

import (
	"fmt"
	"go/ch2/tempconv"
)

type Kelvin float64

const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func C2K(c tempconv.Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func K2C(k Kelvin) tempconv.Celsius {
	return tempconv.Celsius(k - 273.15)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", k)
}

func main() {
	fmt.Println(C2K(tempconv.AbsoluteZeroC))
	fmt.Println(K2C(BoilingK))
}
