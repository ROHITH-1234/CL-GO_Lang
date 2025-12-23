package main
import(
	"fmt"
)
func loops(){
arr :=[]string{"string","int","bollean","float"}
fmt.Println("while loop ")
i:=0
for i <len(arr){
fmt.Print(arr[i]," ")
i++
}
fmt.Println()
fmt.Println("For loop ")
for j:=0;j<len(arr);j++{
	fmt.Print(arr[j]," ")
}
fmt.Println()
for index,value:=range arr{
	fmt.Printf("In index %v has value %v\n",index,value)
}

}