package innova4j

import (
	"fmt"
	"time"
)

func test(from string) {
    for i := 0; i < 10; i++ {
   	 	time.Sleep(100 * time.Millisecond)
        fmt.Println(from, ":", i)
    }
}

func suma(values []int, ch chan int){
	
	defer close(ch)

	sum := 0
	for _, value := range values{
		sum+=value	
	}
	ch <- sum 
}

func Goroutines(){

	go test("goroutine")
	
	test("direct")
	
	 go func(msg string) {
        fmt.Println(msg)
    }("going")

	
	//var input string
    //fmt.Scanln(&input)
    fmt.Println("done")
}


/*Los canales son un tipo de datos a través de los cuales puedes enviar o recibir valores con el operador <-.*/
func Channels(){
	//ch <- v    Envía v al canal ch.
	//v := <-ch   Recibe del canal ch y asigna el valor a v.
	//(Los datos fluyen en la dirección de la "flecha".)
	
	values := []int{3,4,2,-5,1}
	
	channel := make(chan int)
	
	go suma(values[:2], channel)
	//go suma(values[2:], channel)
	
	x, ok := <- channel
	
	y, ok2 := <- channel
	
	fmt.Println(x, "estado ", ok, y, "estado ", ok2)
	/*
	for i := range channel {
		fmt.Println(i)
	}*/
	
	
	
}
