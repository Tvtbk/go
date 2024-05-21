package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	var strA, strB string
	SecondStr := false
	for _, symbol := range text {
		if SecondStr == true {
			strB += fmt.Sprintf("%c", symbol)
		} else if fmt.Sprintf("%c", symbol) == ";" {
			SecondStr = true
			continue
		} else {
			strA += fmt.Sprintf("%c", symbol)
		}
	}
	strA = strings.Replace(strA, ",", ".", -1)
	strB = strings.Replace(strB, ",", ".", -1)
	strA = strings.Replace(strA, " ", "", -1)
	strB = strings.Replace(strB, " ", "", -1)

	flotA, err := strconv.ParseFloat(strA, 64)
	if err != nil {
		panic(err)
	}
	flotB, err := strconv.ParseFloat(strB, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.4f", (flotA / flotB))

}
