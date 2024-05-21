package main

import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(i uint) uint {
		i64 := uint64(i)
		strPerebor := strconv.FormatUint(i64, 10)
		var strI string
		for _, symbol := range strPerebor {
			num := symbol - '0'
			numInt := int(num)
			if int(num) != 0 && int(num)%2 == 0 {
				strI += strconv.FormatInt(int64(numInt), 10)
			} else {
				continue
			}
		}
		if strI == "" {
			return 100
		}
		iRev, err := strconv.Atoi(strI)
		if err != nil {
			panic(err)
		}
		return uint(iRev)
	}
	var d uint = 3456789
	fmt.Println(fn(d))
}
