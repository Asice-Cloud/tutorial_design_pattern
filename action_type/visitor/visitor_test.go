package visitor

import "testing"

func Test(t *testing.T) {
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}

	objectStructure := &ObjectStructure{}
	objectStructure.Attach(elementA)
	objectStructure.Attach(elementB)

	objectStructure.Accept(visitor1)
	objectStructure.Accept(visitor2)
}
