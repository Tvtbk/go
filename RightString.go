package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	textRunes := []rune(text)
	dlen := utf8.RuneCountInString(text)
	dlen--
	index := dlen - 1
	if unicode.IsUpper(textRunes[0]) && fmt.Sprintf("%c", textRunes[index]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
