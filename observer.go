package main

import (
	"container/list"
	"fmt"
)

type Subject interface {
	Attach(Observer)
	Delete(Observer)
	Notify()
}

// 目标-被观察者(容量、载体)
type NomalSubject struct {
	observer *list.List
	value    int
}

// (初始化)实例化目标观察者
func NewSubject() *NomalSubject {
	return &NomalSubject{
		observer: list.New(),
		value:    0,
	}
}

// 添加或绑定观察者
func (s *NomalSubject) Attach(ob Observer) {
	s.observer.PushBack(ob)
}

// 释放或删除观察者
func (s *NomalSubject) Delete(ob Observer) {
	for e := s.observer.Front(); e != nil; e = e.Next() {
		if e.Value.(*Observer) == &ob {
			s.observer.Remove(e)
			break
		}
	}
}

// 通知所有观察者
func (s *NomalSubject) Notify() {
	for e := s.observer.Front(); e != nil; e = e.Next() {
		e.Value.(Observer).Update(s)
	}
}

// 设置目标值
func (s *NomalSubject) SetValue(v int) {
	s.value = v
	s.Notify()
}

func (s *NomalSubject) GetValue() int {
	return s.value
}

// 观察者
type Observer interface {
	Update(Subject)
}

// 观察者类
type FirstObserver struct {
}

//Update 观察者接受被观察目标的更新信息
func (ob *FirstObserver) Update(subject Subject) {
	fmt.Println("first observer : ", subject.(*NomalSubject).GetValue())
}

type SecondObserver struct {
}

func (so SecondObserver) Update(subject Subject) {
	fmt.Println("second observer : ", subject.(*NomalSubject).GetValue())
}

func main() {
	subject := NewSubject()
	observer1 := new(FirstObserver)
	observer2 := new(SecondObserver)
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.SetValue(123456)
}
