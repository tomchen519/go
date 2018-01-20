package main

import "fmt"

// Celsius type float
type Celsius float64

var cel Celsius

// Fahrenheit type float
type Fahrenheit float64

var fah Fahrenheit

// Kelvin type float
type Kelvin float64

var kel Kelvin

// declare constants
const (
	AbsolutZeroCel Celsius = -273.15
	FreezingCel    Celsius = 0
	BoilingCel     Celsius = 100
)

func main() {
	fmt.Println(CelToFah(52))
	fmt.Println(FahToCel(52))
	fmt.Println(CelToKel(52))
	fmt.Println(KelToCel(52))
	fmt.Println(FahToKel(52))
	fmt.Println(KelToFah(52))

	var s = "c"
	val := 64
	if s == "c" {
		cel = Celsius(val)
	} else if s == "f" {
		fah = Fahrenheit(val)
	}
}

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
