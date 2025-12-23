package main

import (
	"os"
)

func savebill(b bill) error {
	data := b.format()
	return os.WriteFile("bill.txt", []byte(data), 0644)
}
