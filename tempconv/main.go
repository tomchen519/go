package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// NEW COMMENTS

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
	// fmt.Println(CelToFah(52))
	// fmt.Println(FahToCel(52))
	// fmt.Println(CelToKel(52))
	// fmt.Println(KelToCel(52))
	// fmt.Println(FahToKel(52))
	// fmt.Println(KelToFah(52))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value to convert: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("A number was not entered")
		return
	}

	fmt.Print("Enter original unit (c, f, or k): ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	var orgUnit string

	if text == "c" || text == "f" || text == "k" {
		orgUnit = text
		orgUnit = strings.TrimSpace(text)
	} else {
		fmt.Println("Incorrect selection for original unit.")
		return
	}

	fmt.Print("Enter convert to unit (c, f, or k): ")
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	var toUnit string

	if text == "c" || text == "f" || text == "k" {
		toUnit = text
		toUnit = strings.TrimSpace(text)
	} else {
		fmt.Println("Incorrect selection for conversion unit.")
		return
	}

	switch orgUnit {
	case "c":
		if toUnit == "f" {
			fmt.Println(CelToFah(Celsius(value)))
		} else if toUnit == "k" {
			fmt.Println(CelToKel(Celsius(value)))
		}
	case "f":
		if toUnit == "c" {
			fmt.Println(FahToCel(Fahrenheit(value)))
		} else if toUnit == "k" {
			fmt.Println(FahToKel(Fahrenheit(value)))
		}
	case "k":
		if toUnit == "c" {
			fmt.Println(KelToCel(Kelvin(value)))
		} else if toUnit == "f" {
			fmt.Println(KelToFah(Kelvin(value)))
		}
	default:
		fmt.Println("Incorrect selection.")
		return
	}

}

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
