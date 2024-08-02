package observer

import "fmt"

/*
观察者模式是一种行为型设计模式，用于在对象之间建立一对多的依赖关系。具体而言，当一个对象（称为主题或者被观察者）发生改变时，
它的所有依赖对象（称为观察者）都会得到通知并自动更新。该模式允许主题和观察者之间的松耦合，从而实现对象间的解耦和灵活性。

2. 观察者模式的特点和优点
观察者模式具有以下特点和优点：

主题和观察者之间是松耦合的，主题不需要知道观察者的具体实现细节。
可以动态地添加和删除观察者，使系统更加灵活。
主题和观察者之间遵循开闭原则，可以独立地进行扩展和重用。
可以实现一对多的依赖关系，一个主题可以有多个观察者。

3. 观察者模式的实际应用场景举例
观察者模式在现实生活中有很多应用场景，例如：

GUI界面中的事件处理机制，如按钮被点击时的处理动作。
股票市场的实时行情推送。
电商平台中的促销活动通知。
*/

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers() {
	for _, o := range s.observers {
		o.Update()
	}
}

type Observer interface {
	Update()
}

type ConcreteObserver_A struct{}

func (o *ConcreteObserver_A) Update() {
	// do something
	fmt.Println("ConcreteObserver_A is updated")
}

type ConcreteObserver_B struct{}

func (o *ConcreteObserver_B) Update() {
	// do something
	fmt.Println("ConcreteObserver_B is updated")
}
