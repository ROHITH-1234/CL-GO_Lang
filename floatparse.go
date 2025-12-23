package main

import (
	"fmt"
	"strconv"
)

func floatparse() {
	str := "09.23"
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(val)
	}
}
