package main

import "fmt"


func main() {
	var ray []int = []int{1, 2, 3, 4} 
	var ray2 [4]int = [4]int{1, 2, 3, 4} 
	var rayT []int = []int{1, 2, 3, 4} 

	rayT = append(rayT, ray2...) 

	fmt.Println(ray)
	fmt.Println(ray2)
	fmt.Println(rayT)

}