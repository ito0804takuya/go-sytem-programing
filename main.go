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

func (g Greeter) Talk() {
	fmt.Printf("My name is %s\n", g.name)
}

func main() {
	// greeter := 
	var talker Talker
	talker = Greeter{name: "jiro"}
  talker.Talk()
}