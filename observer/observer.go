package observer

import (
	"fmt"
	"sync"
)

// 观察者接口
type Observer interface {
	Update(message string)
}

// 主题接口
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(message string)
}

// 具体主题类型
type ConcreteSubject struct {
	observers sync.Map
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers.Store(observer, struct{}{})
}

func (s *ConcreteSubject) Unregister(observer Observer) {
	s.observers.Delete(observer)
}

func (s *ConcreteSubject) Notify(message string) {
	s.observers.Range(func(key, value interface{}) bool {
		observer := key.(Observer)
		observer.Update(message)
		return true
	})
}

// 具体观察者类型
type ConcreteObserver struct {
	Name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s received message: %s\n", o.Name, message)
}
