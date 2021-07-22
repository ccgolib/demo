package design_pattern

import (
	"container/list"
	"fmt"
)

type Subject interface {
	Add(Observer)
	Delete(Observer)
	Notify()
}
type ConcreteSubject struct {
	Observer *list.List
	value    int
}

func NewConcreteSubject() *ConcreteSubject {
	s := new(ConcreteSubject)
	s.Observer = list.New()
	return s
}

// 注册/添加观察者
func (s *ConcreteSubject) Add(observer Observer) {
	s.Observer.PushBack(observer)
}

// 释放/删除观察者
func (s *ConcreteSubject) Delete(observer Observer) {
	for ob := s.Observer.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*Observer) == &observer {
			s.Observer.Remove(ob)
			break
		}
	}
}

// 通知所有观察者
func (s *ConcreteSubject) Notify() {
	for ob := s.Observer.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Observer).Update(s)
	}
}

func (s *ConcreteSubject) setValue(value int) {
	s.value = value
	s.Notify()
}

func (s *ConcreteSubject) getValue() int {
	return s.value
}

type Observer interface {
	Update(Subject)
}

// 具体观察者
type ConcreteObserver1 struct {
}

func (c *ConcreteObserver1) Update(subject Subject) {
	value := subject.(*ConcreteSubject).getValue()
	fmt.Println("ConcreteObserver1  value is ", value)
}

// 具体观察者
type ConcreteObserver2 struct {
}

func (c *ConcreteObserver2) Update(subject Subject) {
	println("ConcreteObserver2 value is ", subject.(*ConcreteSubject).getValue())
}

//func main() {
//	Subject := NewConcreteSubject()
//	observer1 := new(ConcreteObserver1)
//	observer2 := new(ConcreteObserver2)
//	Subject.Add(observer1)
//	Subject.Add(observer2)
//	Subject.setValue(5)
//}
