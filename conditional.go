package main
import(
	"fmt"
)
func conditional() {
	a:=10
	for i:=0;i<=a;i++ {
		if i==a {
			  fmt.Println("Reached the end of the loop",i)
			  break
			  
			} else if i==0 {
			  fmt.Println("dont print 0")
			  continue
		} else {
				fmt.Println(i)
			}
	}
}
