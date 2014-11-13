package main

import "fmt"


func main() {
	a := make([]int, 1)
	// a == []int{0}
	a = append(a, 1, 2, 3)
	// a == []int{0, 1, 2, 3}
	fmt.Printf("%d \n",a[1])

}