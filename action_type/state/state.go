package state

import "fmt"

/*
状态模式是一种行为型设计模式，用于解决对象在不同状态下有不同行为的问题。
它将对象的行为封装到不同的状态类中，使得对象在运行时可以根据其内部状态的变化而改变行为。

2. 状态模式的特点和优点
状态模式的主要特点和优点如下：

将复杂的状态判断逻辑封装到不同的状态类中，增强了代码的可维护性。
开闭原则：通过添加新的状态类，可以方便地增加新的状态。
各个状态类之间相互独立，修改一个状态类不会影响其他状态类的代码。
简化了条件语句的判断逻辑，提高了代码的可读性和可扩展性。
3. 状态模式的实际应用场景举例
状态模式在现实生活中有很多应用场景，例如：

交通信号灯：根据不同的状态，交通信号灯会发出不同的光信号。
订单状态管理：订单在不同的状态下有不同的操作和行为,如已支付、已发货、已签收等
*/

// 状态接口
type State interface {
	Handle(context *Context)
}

// 具体状态类A
type ConcreteStateA struct {
	name string
}

func (c *ConcreteStateA) Handle(context *Context) {
	fmt.Println("当前状态为：", c.name)
	fmt.Println("执行具体状态A的操作...")
}

// 具体状态类B
type ConcreteStateB struct {
	name string
}

func (c *ConcreteStateB) Handle(context *Context) {
	fmt.Println("当前状态为：", c.name)
	fmt.Println("执行具体状态B的操作...")
}

// 上下文类
type Context struct {
	state State
}

// 处理订单
func (c *Context) Handle() {
	c.state.Handle(c)
}

// 设置状态
func (c *Context) SetState(state State) {
	c.state = state
}

// 具体状态类A的切换方法
func (c *ConcreteStateA) SwitchStateB(context *Context) {
	context.SetState(&ConcreteStateB{name: "具体状态B"})
}

// 具体状态类B的切换方法
func (c *ConcreteStateB) SwitchStateA(context *Context) {
	context.SetState(&ConcreteStateA{name: "具体状态A"})
}
