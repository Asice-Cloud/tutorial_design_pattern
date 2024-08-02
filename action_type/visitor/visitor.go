package visitor

import "fmt"

/*
访问者模式是一种行为型设计模式，通过将数据结构与数据操作分离，实现对数据的不同操作而不改变数据结构。
访问者模式可以将数据结构与操作解耦，从而使操作更加灵活和可扩展。

2. 访问者模式的特点和优点
特点：

将数据结构与操作解耦，可实现不同操作的动态绑定。
添加新的操作非常方便，无需修改已有代码。
优点：

增加新的操作非常方便，符合开闭原则。
可以对数据结构进行复杂的操作，而无需改变数据结构本身。
3. 访问者模式的实际应用场景举例
访问者模式在实际应用中具有广泛的应用场景，例如：

编译器的语法树分析阶段，可以使用访问者模式来实现不同的语法检查和代码转换操作。
数据库查询优化器中，可以使用访问者模式来实现对查询树的不同优化操作。
*/

type Visitor interface {
	VisitorConcreteElementA(element ConcreteElementA)
	VisitorConcreteElementB(element ConcreteElementB)
}

type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitorConcreteElementA(element ConcreteElementA) {
	element.OperationA()
}

func (v *ConcreteVisitor1) VisitorConcreteElementB(element ConcreteElementB) {
	element.OperationB()
}

type ConcreteVisitor2 struct{}

func (v *ConcreteVisitor2) VisitorConcreteElementA(element ConcreteElementA) {
	element.OperationA()
}
func (v *ConcreteVisitor2) VisitorConcreteElementB(element ConcreteElementB) {
	element.OperationB()
}

type Element interface {
	Accept(visitor Visitor)
}
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitorConcreteElementA(*e)
}
func (e *ConcreteElementA) OperationA() {
	fmt.Println("ConcreteElementA OperationA")
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitorConcreteElementB(*e)
}
func (e *ConcreteElementB) OperationB() {
	fmt.Println("ConcreteElementB OperationB")
}

type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Attach(element Element) {
	o.elements = append(o.elements, element)
}
func (o *ObjectStructure) Detach(element Element) {
	for i, e := range o.elements {
		if e == element {
			o.elements = append(o.elements[:i], o.elements[i+1:]...)
			break
		}
	}
}
func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, e := range o.elements {
		e.Accept(visitor)
	}
}
