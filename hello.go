// インターフェース
// メソッド一覧を定義したデータ型

package main

import "fmt"

type greeter interface {
	greet()
}

type japanese struct{}
type american struct{}

func show(t interface{}) {
	//// 型アサーション
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("I am Japanese")
	// } else {
	// 	fmt.Println("I am not Japanese")
	// }

	//// 型Switch
	switch t.(type) {
	case japanese:
		fmt.Println("I am Japanese")
	default:
		fmt.Println("I am not Japanese")
	}
}

func (j japanese) greet() {
	fmt.Println("こんにちは")
}

func (a american) greet() {
	fmt.Println("Hello")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	show(japanese{})
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
