package main

import "fmt"

func main() {
	arrays()
	loops()
	multivalue()
	function()
	conditional()
	mapfunc()
	passbyvalue()
	fmt.Println(customtype("Rohith"))

	fmt.Println(customtype("rohith").format())
	k := changes()
	fmt.Println(k.format())

	temp := billcreator()
	switchstatement()
	floatparse()
	fmt.Println(temp.format())
	err := savebill(temp)
	if err != nil {
		fmt.Println("Error saving bill:", err)
	} else {
		fmt.Println("successfull saved the file")
	}

	pointer()
	fmt.Println("hello, Rohith")
	var firstName string = "first"
	var SecoundName = "U"
	fmt.Print(firstName, SecoundName, "\n")
	var a int = 10
	var b int = 20
	fmt.Println(a + b)
	var c int32 = 1234567
	var d int32 = 123456789
	fmt.Printf("%v sum %v", c, d)
	short := "Rohith"
	fmt.Println(short)
}
