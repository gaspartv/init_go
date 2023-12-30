package main

import "fmt"

const ebulicaoF = 212

func main() {

	tempFF := 32
	tempFC := 0
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9 
	fmt.Printf("A temperatura de ebulição da água em °F = %v (%T), temperatura de ebulição da água em °C =%v (%T) .", tempF, tempF, tempC, tempC)
	fmt.Printf("A temperatura de fusão da água em °F = %v (%T), e a temperatura de fusão da água em °C é = %v (%T). ", tempFF, tempFF, tempFC, tempFC)
}