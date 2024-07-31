package factory_mode

import "testing"

func CreateProduct(factory AbstractFactory) {
	productA := factory.CreateProduct_1()
	productB := factory.CreateProduct_2()

	productA.Method_1()
	productB.Method_2()
}

func Test_2(t *testing.T) {
	factory1 := &ConcreteFactory_1{}
	factory2 := &ConcreteFactory_2{}
	CreateProduct(factory1) // Output: Product A1: Method A, Product B1: Method B
	CreateProduct(factory2) // Output: Product A2: Method A, Product B2: Method B
}
