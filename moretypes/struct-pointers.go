package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main()  {
	v := Vertex{1,2}
	p := &v
	p.X = 1e9
	q := &p
	fmt.Println(v)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(*q)
}
