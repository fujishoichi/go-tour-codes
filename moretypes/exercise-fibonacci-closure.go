package main

import "fmt"

func fibonacci() func() int {
	n1 := 0
	n2 := 1
	number := -1
	return func() int {
		number++
		if number < 2 {
			return number
		}
		n := n1 + n2
		n1 = n2
		n2 = n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
