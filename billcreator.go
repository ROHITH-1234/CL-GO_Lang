package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createnewbill() (string, int, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the user name")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	var age int
	fmt.Println("Enter the age")
	fmt.Fscan(os.Stdin, &age)
	var item string
	fmt.Println("Enter the item")
	fmt.Fscan(os.Stdin, &item)
	return name, age, item
}
func billcreator() {
	name, age, item := createnewbill()
	fmt.Println("user name:", name)
	fmt.Println("user age:", age)
	fmt.Println("user item:", item)
}
