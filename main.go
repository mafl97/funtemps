package main

import (
	"flag"
	"fmt"

	"github.com/mafl97/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope".
var fahrenheit float64
var celsius float64
var kelvin float64
var out string
var funfacts string
var t string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene.
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "0.0", "temperaturskala C, K eller F for funfacts")
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje.
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {

	// logikken for flaggene og kall til funksjoner fra conv.

	flag.Parse()

	if isFlagPassed("C") && !isFlagPassed("K") && !isFlagPassed("F") {

	} else if isFlagPassed("K") && !isFlagPassed("C") && !isFlagPassed("F") {

	} else if isFlagPassed("F") && !isFlagPassed("C") && !isFlagPassed("K") {

	} else {
		return
	}

	if isFlagPassed("out") {
		if isFlagPassed("F") {
			if out == "C" {
				fmt.Printf("%.2fF er %.2f°C\n", fahrenheit, conv.FarhenheitToCelsius(fahrenheit))
			}
		} else if isFlagPassed("C") {
			if out == "F" {
				fmt.Printf("%.2f°C er %.2fF\n", celsius, conv.CelsiusToFarhenheit(celsius))
			}
		}
	}

	if isFlagPassed("out") {
		if isFlagPassed("C") {
			if out == "K" {
				fmt.Printf("%.2f°C er %.2fK\n", celsius, conv.CelsiusToKelvin(celsius))
			}
		} else if isFlagPassed("K") {
			if out == "C" {
				fmt.Printf("%.2fK er %.2f°C\n", kelvin, conv.KelvinToCelsius(kelvin))
			}
		}

	}

	if isFlagPassed("out") {
		if isFlagPassed("F") {
			if out == "K" {
				fmt.Printf("%.2fF er %.2fK\n", fahrenheit, conv.FarhenheitToKelvin(fahrenheit))
			}
		} else if isFlagPassed("K") {
			if out == "F" {
				fmt.Printf("%.2fK er %.2fF\n", kelvin, conv.KelvinToFarhenheit(kelvin))
			}
		}

	}

}
