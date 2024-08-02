package template

import "testing"

func Test(t *testing.T) {
	class1 := NewConcreteClass_1()
	class1.TemplateMethod()

	class2 := NewConcreteClass_2()
	class2.TemplateMethod()
}
