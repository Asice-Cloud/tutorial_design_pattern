package memento

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("State 1")
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("State 2")
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetMemento(caretaker.GetMemento(0))
	fmt.Println("Originator state after restoring to state 1:", originator.state)

	originator.SetMemento(caretaker.GetMemento(1))
	fmt.Println("Originator state after restoring to state 2:", originator.state)
}
