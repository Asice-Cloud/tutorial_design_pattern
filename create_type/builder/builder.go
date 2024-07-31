package builder

import "fmt"

/*
建造者模式是一种创建型设计模式，它允许你按照步骤创建复杂对象。它将对象的构建与其表示分离，可以根据所需的组合配置不同的属性和参数

2. 建造者模式的特点和优点
建造者模式的特点和优点包括：

封装了对象的创建和组装过程，使客户端代码与具体的构建过程分离，更加灵活和可维护。
可以通过不同的建造者组合配置不同的属性和参数，以创建不同的对象。
提高了代码的可读性和可维护性，使得代码更易于理解和扩展。
可以避免在构造函数中使用过多的参数，使得代码更加简洁。

3. 建造者模式的应用场景
建造者模式适用于以下场景：

当需要一步一步构建复杂对象时，可以使用建造者模式。
当对象的构建过程比较复杂，并且存在多种不同的组合配置时，可以使用建造者模式。
当需要创建不同表示的对象时，可以使用建造者模式。
*/

/*
建造者模式与其他设计模式的关系包括：

建造者模式可以与抽象工厂模式结合使用，来创建多个产品系列。
建造者模式可以与单例模式结合使用，来创建单例对象的复杂构建过程。
*/

// Builder 是建造者接口，定义了构建产品各部分的方法。
type Builder interface {
	// setPart_A 构建产品的 Part_A 部分。
	setPart_A()
	// setPart_B 构建产品的 Part_B 部分。
	setPart_B()
	// setPart_C 构建产品的 Part_C 部分。
	setPart_C()
	// getResult 返回构建完成的产品。
	getResult() *Product
}

// ConcreteBuilder 是具体建造者，实现了 Builder 接口。
type ConcreteBuilder struct {
	product *Product
}

// setPart_A 实现了 Builder 接口的方法，构建产品的 Part_A 部分。
func (b *ConcreteBuilder) setPart_A() {
	b.product.Part_A = "A"
}

// setPart_B 实现了 Builder 接口的方法，构建产品的 Part_B 部分。
func (b *ConcreteBuilder) setPart_B() {
	b.product.Part_B = "B"
}

// setPart_C 实现了 Builder 接口的方法，构建产品的 Part_C 部分。
func (b *ConcreteBuilder) setPart_C() {
	b.product.Part_C = "C"
}

// getResult 实现了 Builder 接口的方法，返回构建完成的产品。
func (b *ConcreteBuilder) getResult() *Product {
	return b.product
}

// Product 是产品类，包含产品的各个部分。
type Product struct {
	Part_A string
	Part_B string
	Part_C string
}

// Show 显示产品的各个部分。
func (p *Product) Show() {
	fmt.Println(p.Part_A + p.Part_B + p.Part_C)
}

// Director 是指挥者类，负责使用 Builder 构建产品。
type Director struct {
	builder Builder
}

// Construct 构建产品，按照顺序调用 Builder 的方法。
func (d *Director) Construct() {
	d.builder.setPart_A()
	d.builder.setPart_B()
	d.builder.setPart_C()
}

// NewConcreteBuilder 创建并返回一个新的 ConcreteBuilder 实例。
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}
