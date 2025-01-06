package main

import "fmt"

func main() {
	const pi = 3.14159
	var radio float64 = 5.0
	area := pi * radio * radio
	fmt.Printf("El área del círculo con radio %.2f es %.2f\n", radio,
		area)
}
