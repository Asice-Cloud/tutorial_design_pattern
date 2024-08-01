package command_mode

import "fmt"

/*
命令模式是一种行为设计模式，它允许将请求封装为一个对象，从而使你可以用不同的请求对客户端进行参数化。

1.2 命令模式的作用
命令模式的主要目的是解耦发送者和接收者。通过将请求封装成一个对象，发送者只需要和命令对象进行交互，而不需要直接和接收者交互。

2. 命令模式的特点和优点
命令模式具有以下特点和优点：

将请求封装成对象，使发送者和接收者解耦。
可以将请求排队、记录日志、撤销等操作。
可以在不改变原有代码的情况下扩展新的命令。
3. 命令模式的实际应用场景举例
命令模式在以下场景中适用：

需要将请求和执行命令的对象解耦。
需要支持撤销和重做操作。
需要将一组操作排队执行。
*/

/*
命令模式和策略模式在某种程度上是相似的，它们都将某些行为封装成对象。
不同之处在于，命令模式主要用于将请求封装为对象并实现撤销和执行队列等功能，
而策略模式则主要用于封装一系列的算法，并在运行时根据需要选择其中一个算法进行执行。

命令模式更适合于需要记录日志、记账等操作，而策略模式更适合于业务逻辑的灵活变化。
*/

type ICommand interface {
	Execute()
}

type ConcreteCommand struct {
	receiver IReceiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

type IReceiver interface {
	Action()
}

type Receiver struct{}

func (r *Receiver) Action() {
	// do something
	fmt.Println("Receiver.Action")
}

type Invoker struct {
	command ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
