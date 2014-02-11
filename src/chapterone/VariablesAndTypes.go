/* Constants and Variable= Keywords(const, var)
las constantes son declaradas usando (const) 
y las variables(var), tambien se pueden definir sin estas 
palabrasClaves y permite definir variables sin el tipo, 
por defecto si no se define el 
tipo GO asigna (int, float64, String, bool), las variables 
numericas y de caracter si no son inicializadas Go garantiza 
inicializarlas en 0 y un String vacio.

//This means that every numeric variable is guaranteed to be zero and every string to be empty—
var i int
var x string

Definiendo constantes y variables.
const limit = 512 // constant; type-compatible with any number
const top uint16 = 1421 // constant; type: uint16
start := -19 // variable; inferred type: int
end := int64(9876543210) // variable; type: int64
var i int // variable; value 0; type: int
var debug = false // variable; inferred type: bool
checkResults := true // variable; inferred type: bool
stepSize := 1.5 // variable; inferred type: float64
acronym := "FOSS" // variable; inferred type: string


Enumerations. keyword (const)
const Cyan = 0
const Magenta = 1
const Yellow = 2


Boolean and Comparison Operators
Syntax Description/result
!b 		Logical NOT operator; false if Boolean expression b is true
a || b 	Short-circuit logical OR operator; true if either Boolean expression a or b is true
a && b	Short-circuit logical AND operator; true if both Boolean expressions a and b are true
x < y true 	if expression x is less than expression y
x <= y true if expression x is less than or equal to expression y
x == y true if expression x is equal to expression y
x != y true if expression x is not equal to expression y
x >= y true if expression x is greater than or equal to expression y
x > y true	if expression x is greater than expression y


- The == and != operators can be used to compare two pointers or two interfaces—
or to compare a pointer or interface or reference
- (<, <=, >=, >) may be applied only to numbers and strings
- use such as Less() or Equal(), to comparision of function or methods.
*/
package main

import (
"fmt"
)

const ( 
Uno = iota // iota define un valor 0 y las siguientes incrementa en uno
Dos 
tres
)

const Cuatro = 4

func mains(){
	constantes()
	integers()
	
	var doss string  = "HOLA uno dos"
	
	doss += " tres"
	
	fmt.Println(doss," tamaño:", len(doss));
	
	i:= 2

	switch i { 
		case 0: fmt.Println("cero")
		case 1: fmt.Println("uno")
		case 2: fmt.Println("dos")
		case 3: fmt.Println("tres")
		default: fmt.Println("null")
	}
	
	
	var x [3] int  
	x[0] = 1;
	x[1] = 2;
	x[2] = 3;

	total := 0
	for _, value := range x {
		total += value
	}
	fmt.Println("total for: ", total)
	
	//y := make ([]float64, 5, 10)
	y := []float64 {1,2,3,4}
	w := make ([] float64, 2)
	z := append(y, 7)
	
	//copia a w el arreglo y.
	copy(w, y)
	
	fmt.Println("Slice: ", y[1:4], " Append ", z, " Copy ", w)
	
	
	mapper := make (map[string]int)
	mapper["uno"] = 1
	mapper["dos"] = 2
	mapper["tres"] = 3
	mapper["cuatro"] = 4
	
	delete (mapper, "dos")
	fmt.Println("Map: ", mapper)
	
	if name, ok := mapper["tres"]; ok {
		fmt.Println(name, "it is ", ok)
	}
	
	mapper2 := map[string]string{
		"hola" : "hola",
		"dos" : "dos",
	}
	fmt.Println(mapper2)

	val := []float64{
		4, 4,
	}
	result := average(val)
	fmt.Println("Promedio de ", val, " es: ", result)
	
	inc := 0
	add := func() int{
		inc ++
		return inc
	} 
	fmt.Println("Invoca variable add:", add())
	fmt.Println("Invoca variable add:", add())
	
}

func average(val []float64)float64 {
	var total float64 = 0
	
	for _, v := range val {
		total += v
	}

	return total / float64(len(val))
}


func constantes(){
fmt.Println(Uno)
fmt.Println(Dos)
fmt.Println(tres)
fmt.Println(Cuatro)

}

/*Go provides 11 separate integer types, five signed and five unsigned, plus an integer
type for storing pointers—

Type 	Range
byte 	Synonym for uint8
int 	The int32 or int64 range depending on the implementation
int8	[−128, 127]
int16 	[−32768, 32767]
int32 	[−2147483648, 2147483647]
int64 	[−9 223372036854 775808, 9223372036854 775807]
rune 	Synonym for int32
uint 	The uint32 or uint64 range depending on the implementation
uint8 	[0, 255]
uint16 	[0, 65535]
uint32 	[0, 4 294 967295]
uint64 	[0, 18446744 073709551615]
uintptr An unsigned integer capable of storing a pointer value (advanced)
*/
func integers(){

var x uint
//y:= 123
//u:= 45
x = 12

fmt.Println(^x)//The bitwise complement of x
/*fmt.Println(x %= y) //Sets x to be the remainder of dividing x by y; division by zero causes a runtime panic
fmt.Println(x &= y)// Sets x to the bitwise AND of x and y
fmt.Println(x |= y)// Sets x to the bitwise OR of x and y
fmt.Println(x ^= y)// Sets x to the bitwise XOR of x and y
fmt.Println(x &^= y )//Sets x to the bitwise clear (AND NOT) of x and y
fmt.Println(x >>= u)// Sets x to the result of right-shifting itself by unsigned int u shifts
fmt.Println(x <<= u)// Sets x to the result of left-shifting itself by unsigned int u shifts
fmt.Println(x % y)// The remainder of dividing x by y; division by zero causes a runtime panic
fmt.Println(x & y)// The bitwise AND of x and y
fmt.Println(x | y)// The bitwise OR of x and y
fmt.Println(x ^ y)// The bitwise XOR of x and y
fmt.Println(x &^ y)// The bitwise clear (AND NOT) of x and y
fmt.Println(x << u)// The result of left-shifting x by unsigned int u shifts
fmt.Println(x >> u)// The result of right-shifting x by unsigned int u shifts
*/

}
