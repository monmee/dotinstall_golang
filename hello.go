package main

import "fmt"

func main() {
	s := []int{2, 4, 6}
	// for idx, val := range s {
	// 	fmt.Println(idx, val)
	// }

	//// _: blank修飾子
	// for _, val := range s {
	// 	fmt.Println(val)
	// }

	m := map[string]int{"monmee": 100, "yuto": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
