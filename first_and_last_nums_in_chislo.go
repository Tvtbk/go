package main

import "fmt"

func main() {
	var a, b, g int
	fmt.Scan(&a)

	switch {
	case a > 0 && a <= 9:
		d := &b
		*d = a / 1
	case a > 9 && a <= 99:
		d := &b
		*d = a / 10
	case a > 99 && a <= 999:
		d := &b
		*d = a / 100
	case a > 999 && a <= 9999:
		d := &b
		*d = a / 1000
	default:
		d := &b
		*d = a / 10000
	}
	g = a % 10
	fmt.Println("first num:", b)
	fmt.Println("last num:", g)
}
