package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var sum int
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if s.Text() == "" {
			break
		} else if err != nil {
			panic(err)
		}
		sum += num
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
