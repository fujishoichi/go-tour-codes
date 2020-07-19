package main

import "fmt"

func main() {
	var s []int
	//s = []int{1,2,3}[0:0]
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}