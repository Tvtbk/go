package main

import "fmt"

func main() {
	var vvod float64
	fmt.Scan(&vvod)
	if vvod <= 0 {
		fmt.Printf("число %2.2f не подходит", vvod)
	} else if vvod > 10000 {
		fmt.Printf("%e", vvod)
	} else {
		vvod *= vvod
		fmt.Printf("%.4f", vvod)
	}
}
