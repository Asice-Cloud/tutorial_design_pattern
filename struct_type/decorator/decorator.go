package decorator

/*
装饰器模式是一种结构型设计模式，它允许动态地向一个对象添加额外的功能，而无需修改对象的代码。装饰器模式通过将对象包装在装饰器类中，实现在运行时添加、修饰或修改对象行为的能力。

2. 装饰器模式的特点和优点
装饰器模式的特点和优点包括：

动态扩展对象的功能，无需修改对象代码。
符合开闭原则，可动态添加和移除装饰器。
可以组合多个装饰器，实现嵌套式的功能扩展。
装饰器可以独立于对象被装饰的方式独立变化。
3. 装饰器模式的实际应用场景举例
装饰器模式在实际开发中有许多应用场景，例如：

动态添加日志记录功能
动态添加缓存功能
动态数据校验

*/

/*
5.1. 与继承的比较
装饰器模式与继承相比，它可以动态地添加功能而无需修改已有的代码。而继承则是静态的，需要在编译时确定。

5.2. 与静态代理模式的比较
装饰器模式与静态代理模式都可以实现功能扩展，但装饰器模式更加灵活，可以动态添加和移除功能。

5.3. 与动态代理模式的比较
装饰器模式和动态代理模式都可以在运行时对对象进行功能扩展，但装饰器模式是对单个对象进行装饰，而动态代理模式是代理对象与被代理对象之间的间接访问。
*/

// Component 是组件接口，定义了 Operation 方法。
type Component interface {
	// Operation 执行组件的操作。
	Operation() string
}

// ConcreteComponent 是具体组件，实现了 Component 接口。
type ConcreteComponent struct{}

// Operation 实现了 Component 接口的方法，返回具体组件的操作结果。
func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent and concrete operation"
}

// Decorator 是装饰器类，包含一个 Component 对象。
type Decorator struct {
	component Component
}

// Operation 调用被装饰的组件的 Operation 方法。
func (d *Decorator) Operation() string {
	return d.component.Operation()
}

// ConcreteDecorator_A 是具体装饰器 A，继承了 Decorator 类，添加了额外的状态。
type ConcreteDecorator_A struct {
	Decorator
	added_state string
}

// Operation 实现了 Component 接口的方法，返回具体装饰器 A 的操作结果。
func (d *ConcreteDecorator_A) Operation() string {
	return d.Decorator.Operation() + " and concrete decorator A operation"
}

// ConcreteDecorator_B 是具体装饰器 B，继承了 Decorator 类。
type ConcreteDecorator_B struct {
	Decorator
}

// Operation 实现了 Component 接口的方法，返回具体装饰器 B 的操作结果。
func (d *ConcreteDecorator_B) Operation() string {
	return d.Decorator.Operation() + " and concrete decorator B operation"
}
