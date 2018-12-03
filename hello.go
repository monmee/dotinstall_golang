package main

import "fmt"

func main() {
	// signal := "red"
	// switch signal {
	// case "red":
	// 	fmt.Println("Stop")
	// case "yellow":
	// 	fmt.Println("Caution")
	// case "blue":
	// 	fmt.Println("Go")
	// default:
	// 	fmt.Println("Wrong signal")
	// }

	score := 82
	switch {
	case score > 80:
		fmt.Println("Great!")
	case score > 60:
		fmt.Println("Good.")
	default:
		fmt.Println("So so.")
	}
}
