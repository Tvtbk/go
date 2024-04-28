package main

import (
	"fmt"
)

var voice_status int

func voice(status bool) {

	v_m := &voice_status

	switch status {

	case true:
		*v_m = 1000

	default:
		*v_m = 0
	}
	fmt.Println(*v_m)
}

func main() {
	voice(true)
	fmt.Println(voice_status)
}
