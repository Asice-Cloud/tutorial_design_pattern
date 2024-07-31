package prototype

/*
在软件开发中，原型模式是一种创建型设计模式，它允许通过复制现有对象来创建新对象，而无需使用new操作符。这种方式利用了对象之间的克隆关系，将对象的创建和使用分离开来。

2. 原型模式的特点和优点
特点：
允许对象在运行时动态创建。
可以减少对象的创建时间，提高系统性能。
对象的创建和使用相分离，便于管理和扩展。
优点：
提高对象的创建效率。
简化对象的创建过程。
可以动态地增加或减少对象。

3. Golang原型模式的应用场景
原型模式适用于以下场景：

对象的创建比较复杂，但是通过复制已有对象的方式可以更高效地创建对象。
需要动态地创建或克隆对象，而不是通过直接创建新的对象实例。
*/

// Prototype 是一个接口，定义了 clone 方法，用于克隆对象。
type Prototype interface {
	// clone 创建并返回当前对象的一个副本。
	clone() Prototype
}

// PrototypeManager 是一个原型管理器，负责注册和克隆原型对象。
type PrototypeManager struct {
	prototypes map[string]Prototype
}

// NewPrototypeManager 创建并返回一个新的 PrototypeManager 实例。
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Prototype)}
}

// Register 注册一个原型对象。
// 参数:
// - name: 原型对象的名称
// - prototype: 要注册的原型对象
// 返回值: 注册的原型对象
func (pm *PrototypeManager) Register(name string, prototype Prototype) Prototype {
	pm.prototypes[name] = prototype
	return prototype
}

// Unregister 注销一个原型对象。
// 参数:
// - name: 要注销的原型对象的名称
func (pm *PrototypeManager) Unregister(name string) {
	delete(pm.prototypes, name)
}

// Clone 克隆一个原型对象。
// 参数:
// - name: 要克隆的原型对象的名称
// 返回值: 克隆的原型对象，如果名称不存在则返回 nil
func (pm *PrototypeManager) Clone(name string) Prototype {
	prototype, ok := pm.prototypes[name]
	if ok {
		return prototype.clone()
	}
	return nil
}

// Product 是一个具体的原型对象，实现了 Prototype 接口。
type Product struct {
	name string
}

// clone 创建并返回当前 Product 对象的一个副本。
func (p *Product) clone() Prototype {
	return &Product{name: p.name}
}
