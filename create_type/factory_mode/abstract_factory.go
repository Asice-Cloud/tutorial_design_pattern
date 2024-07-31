package factory_mode

import "fmt"

/*
抽象工厂模式是一种创建型设计模式，它提供了一种创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类
*/
/*
抽象工厂模式具有以下特点和优点：

将对象的创建与使用分离，客户端只需关心接口而不需要关心具体实现。
可以遵循单一职责原则和开闭原则，易于扩展和维护。
提供了一种统一的方式来创建一系列产品，让产品之间的关联更加灵活。
*/
/*
抽象工厂模式通常适用于以下场景：
当系统需要一系列相互关联或相互依赖的对象时，可以使用抽象工厂模式来统一创建。
需要一个产品系列的对象，而不关心具体实现时，可以使用抽象工厂模式来创建该系列的对象。
*/

// AbstractFactory 是抽象工厂接口，定义了创建产品的方法。
type AbstractFactory interface {
	// CreateProduct_1 创建产品1的抽象方法。
	CreateProduct_1() AbstractProduct_1
	// CreateProduct_2 创建产品2的抽象方法。
	CreateProduct_2() AbstractProduct_2
}

// ConcreteFactory_1 是具体工厂1，实现了 AbstractFactory 接口。
type ConcreteFactory_1 struct{}

// CreateProduct_1 创建并返回 Product_1_1 的实例。
func (f *ConcreteFactory_1) CreateProduct_1() AbstractProduct_1 {
	return &Product_1_1{}
}

// CreateProduct_2 创建并返回 Product_1_2 的实例。
func (f *ConcreteFactory_1) CreateProduct_2() AbstractProduct_2 {
	return &Product_1_2{}
}

// ConcreteFactory_2 是具体工厂2，实现了 AbstractFactory 接口。
type ConcreteFactory_2 struct{}

// CreateProduct_1 创建并返回 Product_2_1 的实例。
func (f *ConcreteFactory_2) CreateProduct_1() AbstractProduct_1 {
	return &Product_2_1{}
}

// CreateProduct_2 创建并返回 Product_2_2 的实例。
func (f *ConcreteFactory_2) CreateProduct_2() AbstractProduct_2 {
	return &Product_2_2{}
}

// AbstractProduct_1 是产品1的抽象接口，定义了产品1的方法。
type AbstractProduct_1 interface {
	// Method_1 是产品1的方法。
	Method_1()
}

// AbstractProduct_2 是产品2的抽象接口，定义了产品2的方法。
type AbstractProduct_2 interface {
	// Method_2 是产品2的方法。
	Method_2()
}

// Product_1_1 是产品1的具体实现，包含 Method_1 方法。
type Product_1_1 struct{}

// Method_1 是 Product_1_1 的具体实现，打印产品1的方法。
func (p *Product_1_1) Method_1() {
	fmt.Println("Product_1_1, Method_1")
}

// Product_1_2 是产品2的具体实现，包含 Method_2 方法。
type Product_1_2 struct{}

// Method_2 是 Product_1_2 的具体实现，打印产品2的方法。
func (p *Product_1_2) Method_2() {
	fmt.Println("Product_1_2, Method_2")
}

// Product_2_1 是产品1的具体实现，包含 Method_1 方法。
type Product_2_1 struct{}

// Method_1 是 Product_2_1 的具体实现，打印产品1的方法。
func (p *Product_2_1) Method_1() {
	fmt.Println("Product_2_1, Method_1")
}

// Product_2_2 是产品2的具体实现，包含 Method_2 方法。
type Product_2_2 struct{}

// Method_2 是 Product_2_2 的具体实现，打印产品2的方法。
func (p *Product_2_2) Method_2() {
	fmt.Println("Product_2_2, Method_2")
}
