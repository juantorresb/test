package innova4j

import (
	"fmt"
)

func Slices(){
	
	//inicializar slices con make
	var slice  = make([]string, 5)
	
	// se imprime el slice y el tamaño 
 	fmt.Println(slice, len(slice))
 	
 	//Asignacion de elementos al slice.
 	slice[0] = "Hello"
 	slice[2] = "Innova4j"
 	slice[4] = "Team"
 	
 	//se obtienen posiciones especificas.
 	fmt.Println(slice[0], slice[3])	
 	
 	slice2 := make([]string, len(slice))

 	//copia los elementos en slice2 (2) del elemento slice(1) 	
 	copy(slice2, slice)
 	fmt.Println("Copy", slice2)
 	
 	x := slice2[:5]
 	fmt.Println("Obtiene los elementos del 0 al 5", x)
 	
 	y := slice2[3:]
 	fmt.Println("Obtiene los elementos del 3 al 5", y)
 	
 	l := slice2[2:5]
 	fmt.Println("Obtiene los elementos del 2 al 5", l)
 	
 	//declarar un arreglo sin tamaño 
 	arregloSinTamano := []string {"1", "2", "3",}
 	//arregloSinTamano[3] := "123" // no se puede definir una nueva posicion.
 	fmt.Println(arregloSinTamano)
 	
 	
 	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen) //Se inicializa el arreglo en la posicion i.
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
 	
 	
 	
 	
 	
 	
}

