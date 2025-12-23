package main
import(
	"fmt"
)
func operation(a []int) (int, int, int){
	add:=0
 for _,value:=range a {
add+=value
 }
 mul:=1
 for _,value:=range a {
mul*=value
 }
 sub:=0
 for _,value:=range a {
sub*=value
 }
 return add,sub,mul
}
func multivalue(){
  ages:=[]int{10,20,30,40}
  add,sub,mul:=operation(ages)
  fmt.Println(add)
  fmt.Println(sub)
  fmt.Println(mul)
}