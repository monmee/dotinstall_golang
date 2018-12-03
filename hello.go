package main

import "fmt"

func main() {
	// var a [5]int //a[0]-a[4]
	// a[2] = 1
	// a[3] = 2
	// fmt.Println(a)
	// b := [3]int{1, 3, 5}
	b := [...]int{1, 3, 5}
	fmt.Println(b)
	fmt.Println(len(b))
}
