package main

import "fmt"

func main() {
	// m := make(map[string]int)
	// m["monmee"] = 200
	// m["yuto"] = 100
	m := map[string]int{"monmee": 200, "yuto": 100}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "monmee")
	fmt.Println(m)
	val, isOk := m["yuto"]
	fmt.Println(val)
	fmt.Println(isOk)
}
