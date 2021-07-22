package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type cat struct{}

func (c cat) say() {
	fmt.Println("cat!cat!")
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

type car struct {
	brand string
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main() {
	//var s Sayer
	//a := cat{}
	//b := dog{}
	//s = a
	//s.say()
	//s = b
	//s.say()

	//var m Mover
	//var wangcao = dog{}
	//m = wangcao
	//m.move()
	//var fugui = &dog{}
	//m = fugui
	//m.move()

	var s Sayer
	var m Mover
	var a = dog{name: "旺财"}
	var b = car{brand: "保时捷"}

	s = a
	m = a
	s.say()
	m.move()
	m = b
	m.move()

}
