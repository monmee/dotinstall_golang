package main

import "fmt"

func main() {
	// const name = "monmee"
	// name = "tanaka"
	// fmt.Printf(name)

	const (
		sun = iota // 0
		mon // 1
		tue //2
	)
	fmt.Println(sun, mon, tue)

}
