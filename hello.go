package main

import "fmt"

type user struct {
	name string
	score int
}
func main() {
	//// ポインタが返る
	// u := new(user)
	// (*u).name = "monmee"
	// u.name = "monmee"
	// u.score = 39

	//// ポインタは返らない
	// u := user{"monmee", 100}
	u := user{name: "monmee", score: 100}
	fmt.Println(u)
}
