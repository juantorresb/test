package innova4j

import (
	"fmt"
)

/*
func Select(){

}*/


func Switch(c, quit chan int) {
 x, y := 0, 1
	for {
		select {
			case c <- x:
	            x, y = y, x+y
	        case <-quit:
	            fmt.Println("sale")
	            return
	        }
	 }
}

