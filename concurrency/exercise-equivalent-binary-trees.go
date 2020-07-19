package main
// https://tour.golang.org/concurrency/7

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_Walk(t, ch)
	close(ch)
}

func _Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		_Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		_Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	t1Values := make([]int, 0, 0)
	t2Values := make([]int, 0, 0)
	for v1 := range ch1 {
		t1Values = append(t1Values, v1)
	}
	for v2 := range ch2 {
		t2Values = append(t2Values, v2)
	}
	if len(t1Values) != len(t2Values) {
		return false
	}

	for i := 0; i < len(t1Values); i++ {
		if t1Values[i] != t2Values[i] {
			return false
		}
	}
	return true
}

func testWalk() {
	ch := make(chan int, 1)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	testWalk()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}