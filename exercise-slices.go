package main

import "fmt"

func main() {
	var dy int = 6
	var dx int = 5
	a := make([][]int, dy)
	fmt.Println(a)
	for i, _ := range a {
		a[i] = make([]int, dx)
	}
	fmt.Println(a)

	for i, _ := range a {
		for j, _ := range a[i] {
			a[i][j] = i * j
		}
	}
	fmt.Println(a)
}
