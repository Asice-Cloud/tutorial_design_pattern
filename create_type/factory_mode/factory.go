package factory_mode

/*
工厂方法模式的特点包括：

将对象的创建和使用解耦，客户端只需关心工厂方法和抽象产品接口。
可以通过添加新的具体工厂类和具体产品类来扩展系统，符合开闭原则。
工厂方法模式的优点包括：

封装了对象的创建过程，使系统更加灵活和可扩展。
隐藏了具体产品类的实现细节，减少了客户端对具体类的依赖和耦合。
提供了一种标准化的产品创建方式，方便了系统的维护和扩展。

*/
/*
工厂方法模式适用于以下场景：

客户端不依赖具体产品类，而是依赖于抽象产品接口。
客户端需要根据不同的条件动态创建不同的产品对象。
需要对产品对象的创建过程进行封装和隐藏。
*/

// Product 是产品接口，定义了产品的操作方法。
type Product interface {
	// Operation 执行产品的操作。
	Operation() string
}

// ConcreteProduct_A 是具体产品A，实现了 Product 接口。
type ConcreteProduct_A struct{}

// Operation 实现了 Product 接口的方法，返回具体产品A的操作结果。
func (c *ConcreteProduct_A) Operation() string {
	return "ConcreteProduct_A"
}

// ConcreteProduct_B 是具体产品B，实现了 Product 接口。
type ConcreteProduct_B struct{}

// Operation 实现了 Product 接口的方法，返回具体产品B的操作结果。
func (c *ConcreteProduct_B) Operation() string {
	return "ConcreteProduct_B"
}

// Factory 是工厂接口，定义了创建产品的方法。
type Factory interface {
	// CreateProduct 创建并返回一个产品实例。
	CreateProduct() Product
}

// ConcreteFactory_A 是具体工厂A，实现了 Factory 接口。
type ConcreteFactory_A struct{}

// CreateProduct 创建并返回 ConcreteProduct_A 的实例。
func (c *ConcreteFactory_A) CreateProduct() Product {
	return &ConcreteProduct_A{}
}

// ConcreteFactory_B 是具体工厂B，实现了 Factory 接口。
type ConcreteFactory_B struct{}

// CreateProduct 创建并返回 ConcreteProduct_B 的实例。
func (c *ConcreteFactory_B) CreateProduct() Product {
	return &ConcreteProduct_B{}
}
