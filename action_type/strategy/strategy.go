package strategy

/*
策略模式是一种行为型设计模式，它允许我们根据不同的情况选择不同的算法或行为。
它将不同的算法封装在独立的策略类中，并使这些策略类可以互相替换。通过使用策略模式，我们可以在运行时动态地改变一个对象的行为，而无需直接修改对象的结构。

2. 策略模式的特点和优点
策略模式有以下几个特点和优点：

策略类可以独立变化，增加新的策略类对原有代码没有影响，符合开闭原则。
客户端可以根据需要选择不同的策略，符合单一职责原则。
策略模式提供了可重用的算法或行为，可以避免代码重复。
策略模式提供了一种更好的组织代码的方式，使代码更加清晰和易于维护。
3. 策略模式的实际应用场景举例
策略模式广泛应用于以下场景：

不同的支付方式，例如支付宝、微信支付等。
不同的排序算法，例如冒泡排序、快速排序等。
不同的日志记录方式，例如控制台输出、文件输出等。
*/

type Strategy interface {
	Execute(data interface{}) string
}

type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute(data interface{}) string {
	return "ConcreteStrategyA"
}

type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute(data interface{}) string {
	return "ConcreteStrategyB" + data.(string)
}

/*
实现上下文类Context，该类封装了具体的策略对象，并提供SetStrategy(strategy Strategy)方法用于设置策略对象，
以及ExecuteStrategy(data interface{}) string方法用于执行具体的策略
*/
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data interface{}) string {
	if c.strategy == nil {
		return "there is no strategy"
	} else {
		return c.strategy.Execute(data)
	}
}
