package innova4j

import (
	"fmt"
	"os"
)


func Files(){
	
	buf := make([]byte, 1024)
	fmt.Println("making buf...")
	
	f, hasError := os.Open("/devel/pwd.txt")
	
	fmt.Println("error?", hasError) 
	
	defer f.Close()
	 
	i := 0
	
	for {
		n, _ := f.ReadAt(buf, int64(i)) 
		//fmt.Println("count ", n) 
		
		//
		
		if n == 0 {
			break 
		}
		if(i < n){
			os.Stdout.Write(buf[i:n]) 
			
		}
		i += 10
		
	}
	
}

