package observer_test

import (
	"dataStructuresAndAlgorithms/observer"
	"testing"
)

func TestObserver(t *testing.T) {
	subject := &observer.ConcreteSubject{}

	observer1 := &observer.ConcreteObserver{Name: "Observer 1"}
	observer2 := &observer.ConcreteObserver{Name: "Observer 2"}
	observer3 := &observer.ConcreteObserver{Name: "Observer 3"}

	subject.Register(observer1)
	subject.Register(observer2)
	subject.Register(observer3)

	subject.Notify("Hello, observers!")

	subject.Unregister(observer2)

	subject.Notify("Goodbye, observer 2!")
}
