package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var a string = "%^80"
	var b string = "hhhhh20&&&&nd"
	fmt.Println(adding(a, b))
}
func adding(a, b string) int64 {
	var revA string
	for _, symbol := range a {
		if unicode.IsDigit(symbol) {
			revA += string(symbol)
		}
	}
	var revB string
	for _, symbol := range b {
		if unicode.IsDigit(symbol) {
			revB += string(symbol)
		}
	}
	revAInt, err := strconv.ParseInt(revA, 10, 64)
	if err != nil {
		panic(err)
	}
	revBInt, err := strconv.ParseInt(revB, 10, 64)
	if err != nil {
		panic(err)
	}
	result := revAInt + revBInt
	return result
}
