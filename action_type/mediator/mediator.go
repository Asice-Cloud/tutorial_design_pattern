package mediator

/*
中介者模式是一种行为型设计模式，它通过将对象之间的通信转移给一个中介者对象，从而减少对象之间的直接依赖关系。
在中介者模式中，各个对象不再直接与彼此通信，而是通过中介者对象进行通信。

2. 中介者模式的特点和优点
中介者模式的特点和优点如下：

减少了对象之间的直接耦合，降低了系统复杂性。
可以简化对象之间的通信，由中介者对象统一处理对象之间的通信。
可以集中控制对象之间的交互，便于扩展和维护。
3. 中介者模式的实际应用场景举例
中介者模式在现实生活中有许多应用场景。例如，一个机场调度系统中，调度员充当中介者角色，飞机、地面交通等各个模块作为同事类，通过调度员进行协调和通信。
*/

type Mediator interface {
	Register(name string, colleague Colleague)
	SendMessage(message string, colleague Colleague)
}

type ConcreteMediator struct {
	colleagueMap map[string]Colleague
}

func (c *ConcreteMediator) Register(name string, colleague Colleague) {
	c.colleagueMap[name] = colleague
}

func (c *ConcreteMediator) SendMessage(message string, colleague Colleague) {
	for _, v := range c.colleagueMap {
		if v != colleague {
			v.ReceiveMessage(message)
		}
	}
}

type Colleague interface {
	ReceiveMessage(message string)
	SendMessage(message string)
	GetName() string
}

type ConcreteColleague_A struct {
	mediator Mediator
	name     string
}

func (c *ConcreteColleague_A) ReceiveMessage(message string) {
	println("ConcreteColleague_A receive message: ", message)
}

func (c *ConcreteColleague_A) SendMessage(message string) {
	c.mediator.SendMessage(message, c)
}

func (c *ConcreteColleague_A) GetName() string {
	return c.name
}

type ConcreteColleague_B struct {
	mediator Mediator
	name     string
}

func (c *ConcreteColleague_B) ReceiveMessage(message string) {
	println("ConcreteColleague_B receive message: ", message)
}
func (c *ConcreteColleague_B) SendMessage(message string) {
	c.mediator.SendMessage(message, c)
}
func (c *ConcreteColleague_B) GetName() string {
	return c.name
}
