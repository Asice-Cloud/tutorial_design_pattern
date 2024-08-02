package observer

import (
	"testing"
)

func Test(t *testing.T) {
	subject := &ConcreteSubject{}
	observerA := &ConcreteObserver_A{}
	observerB := &ConcreteObserver_B{}
	subject.RegisterObserver(observerA)
	subject.RegisterObserver(observerB)
	subject.NotifyObservers()
	subject.RemoveObserver(observerA)
	subject.NotifyObservers()
}
