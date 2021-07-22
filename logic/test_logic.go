package logic

import "fmt"

type PersonLogic struct {
	GetAbc func()
}

func (pg *PersonLogic) GetName(n string) string {
	return "my name is " + n
}

func (pg PersonLogic) GetAge(a int) string {
	return fmt.Sprintf("my age is %d", a)
}