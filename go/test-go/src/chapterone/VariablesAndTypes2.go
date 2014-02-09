package main

import (
 	"fmt"
)

func main(){
	par := obtenerPares()
	fmt.Println(par())
	fmt.Println(par())
	
	
	deferPanicAndRecover()
}

func obtenerPares() func() uint {
 	x := uint(0)
  
  	return func() (ret uint){
  		ret = x
  		x += 2
  		
  		return 
	}  	
}

func deferPanicAndRecover(){
	
	/*defer func(){
		str:=recover()
		
	}*/
	//panic("ERROR")	
}
