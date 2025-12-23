package main
import (
	"fmt"
)
func sum(a int,b int) int{
	return a+b
}
func display(name []string) {
   for i:=0;i<len(name);i++ {
		fmt.Print(name[i]," ")
	}
}
func greeting() {
	fmt.Print("Hello DataType ")
}
func welcomedisplay(name []string,f func()) {
for _,value:=range name {
	f()
	fmt.Println(value)
}
}
func function(){
	 a:=5
	b:=5
	fmt.Println("sum of two variable",sum(a,b))
    name:=[]string{"int","boolean","float"}
	 display(name)
	 greeting()
	 fmt.Println()
	 welcomedisplay(name,greeting)
}