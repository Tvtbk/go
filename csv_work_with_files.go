package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)
	for i := 1; i <= 3; i++ {
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil {
			panic(err)
		}
	}
	w.Flush()
	r := csv.NewReader(buf)
	data, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range data {
		fmt.Println(row)
	}
}
