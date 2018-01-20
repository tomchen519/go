package main

// CelToFah converts celsius to fahrenheit
func CelToFah(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FahToCel converts celsius to fahrenheit
func FahToCel(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CelToKel converts celsius to kelvin
func CelToKel(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KelToCel converts kelvin to celsius
func KelToCel(k Kelvin) Celsius { return Celsius(k - 273.15) }

// FahToKel converts fahrenheit to kelvin
func FahToKel(f Fahrenheit) Kelvin { return Kelvin(FahToCel(f) + 273.15) }

// KelToFah converts kelvin to fahrenheit
func KelToFah(k Kelvin) Fahrenheit { return Fahrenheit(CelToFah(Celsius(k - 273.15))) }
