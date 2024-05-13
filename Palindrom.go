package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	stroka, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	stroka = strings.TrimSuffix(stroka, "\n")
	stroka = strings.TrimSuffix(stroka, "\r")

	strokaRunes := []rune(stroka)

	dlen := utf8.RuneCountInString(stroka)

	var rstroka string
	for i := dlen - 1; i >= 0; i-- {
		rbykva := fmt.Sprintf("%c", strokaRunes[i])
		rstroka += rbykva
	}
	rstrokaRunes := []rune(rstroka)

	//rstrokaRunes = append(rstrokaRunes, 13)
	//fmt.Println(strokaRunes)
	//fmt.Println(rstrokaRunes)

	rstroka = string(rstrokaRunes)
	stroka = string(strokaRunes)

	if stroka == rstroka {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
