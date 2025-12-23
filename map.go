package main
import(
	"fmt"
)
func mapfunc() {
 menu:=map[string]int32{
	"Ram":12,
	"Ravi":36,
 }
 menu["Ram"]=26
 fmt.Println(menu)
 for k,v:=range menu {
	 fmt.Printf("%q age is %v\n",k,v)
	}
	fmt.Println("After Deleting")
	delete(menu,"Ravi")
 fmt.Println(menu["Ram"])
 // example program count letters in string using map
 s:="counting the letter"
 count:=map[string]int{}
 for i:=0;i<len(s);i++ {
	if s[i]!=' ' {
		count[string(s[i])]++
	}
 }
 fmt.Println(count)
}


