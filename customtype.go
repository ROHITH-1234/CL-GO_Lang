package main

import (
	"fmt"
)

type bill struct {
	name string
	item map[string]int64
	tip  float64
}

func customtype(name string) bill {
	b := bill{
		name: "Rohith",
		item: map[string]int64{
			"Tea":    16,
			"orange": 250,
		},
		tip: 12.5,
	}
	return b
}
func (b bill) format() string {
	mesg := "breakdown the item \n"
	for k, v := range b.item {
		mesg += fmt.Sprintf("%s : %d \n", k, v)
	}
	mesg += fmt.Sprintf("tip: %.2f\n", b.tip)
	return mesg
}
