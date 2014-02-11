package main

import (
	"fmt"
)

func main(){

	//Go’s Keywords
	/* break 	default 	func 	interface 	select
	   case 	defer 		go 		map 		struct
	   chan 	else 		goto 	package 	switch
	   const 	fallthrough if 		range 		type
	   continue for 		import 	return 		var
	*/
	
	

	//Predefined Identifiers
	// - Identifiers are case-sensitive.
	// - (including A–Z), are considered to be public—exported in Go terminology
	//   while all others are considered to be private—unexported in Go terminology
	
	/*
		append 	copy 	int8 nil true
		bool 	delete 	int16 panic uint
		byte	error 	int32 print uint8
		cap 	false	int64 println uint16
		close 	float32 iota real uint32
		complex float64 len recover uint64
		complex64 imag 	make rune uintptr
		complex128 int 	new string
	*/
	var append int;
	append = 123
	
	//The blank identifier, _, serves as a placeholder for where a variable is expected in an assignment,
	_ = test()
	
	
	const limit = 512 // constant; type-compatible with any number
	const top uint16 = 1421 // constant; type: uint16
	start := -19 // variable; inferred type: int
	end := int64(9876543210) // variable; type: int64
	var i int // variable; value 0; type: int
	var isCheck bool // variable; value false; type: bool
	var debug = false // variable; inferred type: bool
	checkResults := true // variable; inferred type: bool
	stepSize := 1.5 // variable; inferred type: float64
	acronym := "FOSS" // variable; inferred type: string
	
	
	fmt.Println(append, start, end, i, isCheck, debug, checkResults, stepSize, acronym)
}


func test() int{
	return 123
}

