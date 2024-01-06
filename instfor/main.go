package main

import "fmt"

func main() {
	a := 1
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	for b := 1; b <= 10; b++ {
		fmt.Println(b)
	}

	for c := 1; c <= 10; c++ {
		if c%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for d := 1; d <= 10; d++ {
		switch d {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("Um")
			case 2: fmt.Println("Dois")
			case 3: fmt.Println("Tres")
			case 4: fmt.Println("Quatro")
			case 5: fmt.Println("Cinco")
			case 6: fmt.Println("Seis")
			case 7: fmt.Println("Sete")
			case 8: fmt.Println("Oito")
			case 9: fmt.Println("Nove")
			case 10: fmt.Println("Dez")
		}
	}

	for e := 0;	e < 10; e++ {
		fmt.Println(e)
	}
}
