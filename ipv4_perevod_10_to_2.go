package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ipv4_addr string
	//ip_mas := [9]int{}
	fmt.Printf("\r Input ipv4 address: ")
	fmt.Scan(&ipv4_addr)

	var p_1 string
	var p_2 string
	var p_3 string
	var p_4 string

	var tochky string

	for _, chislo := range ipv4_addr {

		schet := fmt.Sprintf("%c", chislo)

		if schet != "." && len(tochky) == 0 {
			p_1 += schet
		} else if schet == "." && len(tochky) == 0 {
			tochky += schet
		} else if schet != "." && len(tochky) == 1 {
			p_2 += schet
		} else if schet == "." && len(tochky) == 1 {
			tochky += schet
		} else if schet != "." && len(tochky) == 2 {
			p_3 += schet
		} else if schet == "." && len(tochky) == 2 {
			tochky += schet
		} else if schet != "." && len(tochky) == 3 {
			p_4 += schet
		} else if schet == "." && len(tochky) == 3 {
			tochky += schet
		}
	}
	ip_1, err_1 := strconv.Atoi(p_1)
	if err_1 != nil {
		fmt.Println("Error!", err_1)
		return
	}

	ip_2, err_2 := strconv.Atoi(p_2)
	if err_2 != nil {
		fmt.Print("Error!", err_2)
		return
	}

	ip_3, err_3 := strconv.Atoi(p_3)
	if err_3 != nil {
		fmt.Print("Error!", err_3)
		return
	}

	ip_4, err_4 := strconv.Atoi(p_4)
	if err_4 != nil {
		fmt.Print("Error!", err_4)
		return
	}
	fmt.Printf("\r Your output: ")
	fmt.Printf("%b.%b.%b.%b\n", ip_1, ip_2, ip_3, ip_4)
	var exit int
	fmt.Printf("\r enter anything to exit")
	fmt.Scan(&exit)
	//g += fmt.Sprintf("%c\n", chislo)
}
