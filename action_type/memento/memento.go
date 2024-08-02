package memento

/*
备忘录模式是一种行为型设计模式，用于保存和恢复对象的内部状态。它将对象的状态保存到备忘录对象中，以后可以通过备忘录对象将对象恢复到之前的状态。

2. 备忘录模式的特点和优点
备忘录模式的特点和优点包括：

可以在不违反封装原则的情况下保存和恢复对象的内部状态。
可以灵活地管理对象的历史状态，方便进行撤销和重做操作。
可以将状态保存到外部，避免对象内部状态的暴露。
3. 备忘录模式的实际应用场景举例
备忘录模式在实际应用中有很多场景，其中一些例子包括：

文字编辑器中的撤销和重做功能，可以使用备忘录模式保存每次操作的状态。
游戏中的保存和加载功能，可以使用备忘录模式保存游戏进度。
电子邮件客户端中的草稿箱功能，可以使用备忘录模式保存草稿邮件的状态。
*/

type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

// recover from memento
func (o *Originator) SetMemento(m *Memento) {
	o.state = m.GetState()
}

type Caretaker struct {
	Mementos []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.Mementos = append(c.Mementos, m)
}
func (c *Caretaker) GetMemento(index int) *Memento {
	return c.Mementos[index]
}
