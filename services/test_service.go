package services

import (
	"demo/logic"
)

var personLogic logic.PersonLogic

type Person struct {
}


func (p Person)Name(n string) string {
	return personLogic.GetName(n)
}

func (p Person)Age(a int) string {
	return personLogic.GetAge(a)
}