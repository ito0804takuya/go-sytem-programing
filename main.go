package main

import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// 	// fmt.Println("Hello", "World!") // 出力は同じ
// }

// ----------------------------------------------------
type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

// Greeter構造体のメソッド
func (g Greeter) Talk() {
	fmt.Printf("My name is %s\n", g.name)
}

func main() {
	// インターフェースを宣言するための変数
	var talker Talker
	// この時点で、GreeterはTalkerインターフェースを使うことを決める
	talker = Greeter{name: "jiro"}
	talker.Talk()
}
