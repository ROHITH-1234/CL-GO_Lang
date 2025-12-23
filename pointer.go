package main
import(
	"fmt"
)
func memoryref(val* int){
	*val=20
}
func pointer() {
	a:=10
	val:=&a
	fmt.Println(val)
	fmt.Println(a)
	memoryref(val)
	fmt.Println(a)

}