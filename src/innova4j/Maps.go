package innova4j

import (
 	"fmt"
)


func Maps(){
	
	fmt.Println("----------------MAPS --------------")
	mp := make(map[string] int) // mapa clave string y valor numerico.
	fmt.Println(mp)

	//Asignacion de elementos en el map	
	mp["uno"] = 1
	mp["dos"] = 2
	fmt.Println("map:", mp)
	
	//obtener valor de la clave dos.
	value1 := mp["dos"]
	fmt.Println("map[dos]:", value1)
	
	//tamaño de los arreglos.
	fmt.Println("len:", len(mp))
	
	delete(mp, "dos")
	fmt.Println("map:", mp)
	
	value2, containt := mp["dos"]
	fmt.Println("value2:", value2, "lo contiene: ", containt)
	
	map2 := map[string]int {"tres":3, "cuatro":4}
	fmt.Println("map 2:", map2)
	
	
	
}