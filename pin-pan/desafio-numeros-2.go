package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if (i% 3 == 0 && i% 5 == 0) {
			fmt.Printf("Pin-Pan")
			continue
		} else if (i% 5 == 0) {
			fmt.Printf("Pan")
			continue
		} else if (i% 3 == 0) {
			fmt.Println("Pin")
			continue
		} else {
			fmt.Println(i)
			continue
		}
	}
}