package main

import (
	"innova4j"
	"fmt"
)
  
func main(){
 
	innova4j.Slices()
	innova4j.Maps() 
	innova4j.Range() 
	fmt.Println("---return values") 
	fmt.Println(innova4j.MultiValue())
	
	fmt.Println(innova4j.VariadicFunc(4,8))
	
	
}

