package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var passwd string
	fmt.Scan(&passwd)
	lenPasswd := utf8.RuneCountInString(passwd)
	var confirm bool
	if lenPasswd > 5 {
		for _, symbol := range passwd {
			if unicode.IsDigit(symbol) || unicode.Is(unicode.Latin, symbol) {
				confirm = true
				continue
			} else {
				fmt.Println("Wrong password")
				confirm = false
				break
			}

		}
	} else {
		fmt.Println("Wrong password")
	}
	if confirm {
		fmt.Println("Ok")
	}

}
