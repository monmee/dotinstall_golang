package main

import "fmt"

func main() {
	a := [5]int{1, 3, 5, 7, 9}
	s := a[2:4]
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

}
