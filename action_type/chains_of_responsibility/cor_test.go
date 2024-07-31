package chains_of_responsibility

import "testing"

func Test(t *testing.T) {
	handler := &ConcreteHandler_A{}

	// 构建责任链
	handler.SetNext(&ConcreteHandler_B{})

	// 发起请求
	handler.HandleRequest(Request{Condition_1: true, Condition_2: false})
	handler.HandleRequest(Request{Condition_1: false, Condition_2: true})
}
