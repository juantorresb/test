package innova4j

import (
	"fmt"
)

func Range(){
	fmt.Println("----------Range-------")
	//Se define un arreglo.
	arr := []int{1,2,3}
	
	sum := 0
	
	//range se puede utilizar para arreglos y slices.
	
	//con range se recorre el arreglo, range devuelve el indice y el valor del arreglo.
	for _, value := range arr {
		sum += value
	}
	fmt.Println("suma:", sum)
	
	//se obtiene el indice
	for i, value := range arr {
		if value == 3 {
			fmt.Println("index= ", i)
		}
	}	
	
	slice := map[string]string {"uno":"1", "dos":"2"}
	
	for i, value := range slice {
		fmt.Printf("%s -> %s\n", i, value) 
	}
	
	for i, value := range "abcd" {
		fmt.Println(i, value)
	}
	
	
	
}

