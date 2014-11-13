package main

import "fmt"

func main() {
    fmt.Printf("hello world\n")
	
	p := 0
	for i:=0; i < 10; i++ {
		p += i 
	}
	fmt.Printf("%d \n", p)
	
	
	i := 10;
	for i < 15 && i > 0 {
		p += i
		i++
	}
	
	fmt.Println(p)
	
}