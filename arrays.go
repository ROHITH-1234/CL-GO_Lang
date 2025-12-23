package main
import (
	"fmt"
	"strings"
	"sort"
)
func arrays(){
	var ages=[5]int{10,20,30,40,50}
	fmt.Println("Printing the index 1 to 4:",ages[1:4])
	fmt.Println("Printing the index from 0 to 4",ages[:4])
	fmt.Println("Printing the index from 1 to end",ages[1:])
	//slice
	var new_ages=[]int{60,80,70,90}
	new_ages = append(new_ages,100)
	new_ages[1]=65
	fmt.Println("new ages",new_ages)
	sort.Ints(new_ages) 
	fmt.Println("new sorted ages",new_ages)
   var String="hello this is Rohith"
	fmt.Println(strings.Contains(String,"Rohith"))
	fmt.Println(strings.ToUpper(String))
	fmt.Println(strings.ToLower(String))
	fmt.Println(strings.ReplaceAll(String,"Rohith","unknown"))
	fmt.Println(strings.Index(String,"Rohith"))	
}
