package main

import "fmt"

// func swap(a, b int) (int, int) {
// 	return b, a
// }

func main() {
	// fmt.Println(swap(2, 5))

	// f := func(a, b int) (int, int) {
	// 	return b, a
	// }
	// fmt.Println(f(2, 3))

	func(msg string) {
		fmt.Println(msg)
	}("monmee")
}
