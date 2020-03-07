package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < cap(ch1); i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true

}
func main() {
	ch := make(chan int, 10)
	t := tree.New(1)
	go Walk(t, ch)

	for i := 0; i < cap(ch); i++ {
		fmt.Print(<-ch, ", ")
	}

	fmt.Println()

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
