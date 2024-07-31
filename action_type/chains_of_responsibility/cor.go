package chains_of_responsibility

import (
	"errors"
	"fmt"
)

/*
责任链模式是一种行为设计模式，它通过将请求的发送者和接收者解耦，使多个对象都有机会处理这个请求。每个接收者都包含对另一个接收者的引用，
如果自己无法处理该请求，就会将请求转发给下一个接收者，直到请求被处理或者到达链的末尾。

2. 责任链模式的特点和优点
责任链模式的特点和优点如下：

解耦发送者和接收者：发送者无需关心请求是由哪个接收者处理，也不需要知道链中的具体处理者。
灵活性：可以动态地向责任链中添加、移除或者重新排序处理者，无需修改发送者和接收者的代码。
可扩展性：容易扩展责任链，可以方便地添加新的具体处理者。
单一职责原则：每个具体处理者都只需要关心自己的处理逻辑。
可配置性：可以根据需要配置链中的处理者，使得不同的请求有不同的处理者链。
3. 责任链模式的实际应用场景举例
责任链模式在实际应用中有很多场景，例如：

Web应用中的请求处理：可以使用责任链模式来处理不同类型的请求，如身份验证、日志记录、权限验证等。
错误处理：可以使用责任链模式来处理错误，每个处理者负责处理一种类型的错误，并根据需要将错误转发给下一个处理者。
事件处理：可以使用责任链模式来处理不同类型的事件，如用户点击事件、网络请求事件等。

*/

type Handler interface {
	HandleRequest(request Request) error
	SetNext(handler Handler)
}

type Request struct {
	Condition_1 bool
	Condition_2 bool
}

type ConcreteHandler_A struct {
	next Handler
}

func (handler *ConcreteHandler_A) HandleRequest(request Request) error {
	if request.Condition_1 {
		fmt.Println("ConcreteHandler_A handled the request")
		return nil
	} else {
		if handler.next != nil {
			return handler.next.HandleRequest(request)
		}
		return errors.New("no handler available")
	}
}

func (h *ConcreteHandler_A) SetNext(handler Handler) {
	h.next = handler
}

type ConcreteHandler_B struct {
	next Handler
}

func (handler *ConcreteHandler_B) HandleRequest(request Request) error {
	if request.Condition_2 {
		fmt.Println("ConcreteHandler_B handled the request")
		return nil
	} else {
		if handler.next != nil {
			return handler.next.HandleRequest(request)
		}
		return errors.New("no handler available")
	}
}

func (h *ConcreteHandler_B) SetNext(handler Handler) {
	h.next = handler
}
