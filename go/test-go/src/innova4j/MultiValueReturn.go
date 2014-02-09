package innova4j

import (

)

//Multiple Return Values
func MultiValue() (int, int) {

	return 0,1
}

func VariadicFunc(values ... int) (int, int){
	sum:= 0;
	
	for i:=0; i< len(values); i++ {
		sum += values[i]
	}

	sumRange := 0;
	for _, value := range values {
		sumRange += value
	}
	
	return sum, sumRange
}


