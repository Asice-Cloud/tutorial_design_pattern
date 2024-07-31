package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	component := &ConcreteComponent{}

	decorator_a := &ConcreteDecorator_A{}
	decorator_a.component = component

	decorator_b := &ConcreteDecorator_B{}
	decorator_b.component = decorator_a

	result := decorator_b.Operation()
	fmt.Println(result)
}
