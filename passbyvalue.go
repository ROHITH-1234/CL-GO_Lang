package main
import(
	"fmt"
)
func change(a int){
	a=20
}
func changeArr(arr []int) {
	arr[1]=50
}
func passbyvalue() {
	a:=10
	change(a)
	fmt.Println(a)
	arr:=[]int{10,20,30}
	changeArr(arr)
	fmt.Println(arr)
}