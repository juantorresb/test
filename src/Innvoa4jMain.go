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
  
	fmt.Println("---goroutines") 
	innova4j.Goroutines()  
	innova4j.Channels()
	  
	innova4j.Files()
	
	
	arr := []string{"1","2","3"}
	fmt.Print(arr)
	
	arr2 := make([]string, 2)
	fmt.Print(arr2)
	
	var arr3 [3]int
	fmt.Print(arr3)
}    
   
         