package main

import (
	"fmt"
)

func main() {
	var workArray [10]uint8
	for i := 0; i < len(workArray); i++ {
		var g uint8
		fmt.Scan(&g)
		workArray[i] = g
	}
	for i := 1; i <= 3; i++ {
		var a, b uint8
		fmt.Scan(&a, &b)
		s := workArray[a]
		workArray[a] = workArray[b]
		workArray[b] = s
	}
	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i])
		fmt.Print(" ")
	}
}
