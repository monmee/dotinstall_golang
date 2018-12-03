package main

import "fmt"

func main() {

	// score := 43

	if score := 63; score > 80 {
		fmt.Println("Great!")
	} else if score > 60 {
		fmt.Println("Good.")
	} else {
		fmt.Println("So so.")
	}
}
