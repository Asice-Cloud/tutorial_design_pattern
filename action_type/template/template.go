package template

import "fmt"

type AbstractClass interface {
	TemplateMethod()
	AbstractMethod1()
	AbstractMethod2()
	ConcreteMethod()
}

type BaseClass struct {
	AbstractClass
}

func (b *BaseClass) TemplateMethod() {
	b.AbstractMethod1()
	b.AbstractMethod2()
	b.ConcreteMethod()
}
func (b *BaseClass) ConcreteMethod() {
	fmt.Println("ConcreteMethod")
}

type ConcreteClass_1 struct {
	*BaseClass
}

func NewConcreteClass_1() *ConcreteClass_1 {
	class := &ConcreteClass_1{}
	class.BaseClass = &BaseClass{class}
	return class
}

func (c *ConcreteClass_1) AbstractMethod1() {
	fmt.Println("ConcreteClass_1.AbstractMethod_1")
}

func (c *ConcreteClass_1) AbstractMethod2() {
	fmt.Println("ConcreteClass_1.AbstractMethod_2")
}

type ConcreteClass_2 struct {
	*BaseClass
}

func NewConcreteClass_2() *ConcreteClass_2 {
	class := &ConcreteClass_2{}
	class.BaseClass = &BaseClass{class}
	return class
}

func (c *ConcreteClass_2) AbstractMethod_1() {
	fmt.Println("ConcreteClass_2.AbstractMethod_1")
}

func (c *ConcreteClass_2) AbstractMethod_2() {
	fmt.Println("ConcreteClass_2.AbstractMethod_2")
}
