package design_pattern

import "fmt"

type Person interface {
	Eat()
}

type Boy struct {}

type Girl struct {}

func (*Boy)Eat()  {
	fmt.Println("吃米饭。。。")
}
func (*Girl)Eat()  {
	fmt.Println("吃面条....")
}

type My struct {}

func (m *My)Haha(p Person)  {
	p.Eat()
}

func main()  {
	p := &My{}
	p.Haha(new(Boy))
}
