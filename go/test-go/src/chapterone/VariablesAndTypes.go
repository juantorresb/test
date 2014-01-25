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

func main(){
constantes()

}


func constantes(){
fmt.Println(Uno)
fmt.Println(Dos)
fmt.Println(tres)
fmt.Println(Cuatro)

}

