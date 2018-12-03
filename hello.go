package main

import "fmt"

func main() {
	a := 5
	var pa *int 			// int型を格納する領域のアドレスを格納する準備
	pa = &a 					// &a: aのアドレス
	fmt.Println(pa)

	fmt.Println(*pa) 	// *pa: paの領域にあるデータの値 
}
