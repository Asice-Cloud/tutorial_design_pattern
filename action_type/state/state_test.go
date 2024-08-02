package state

import "testing"

func Test(t *testing.T) {
	// 初始化订单上下文
	context := &Context{
		state: &ConcreteStateA{name: "具体状态A"},
	}

	// 处理订单
	context.Handle()

	// 切换状态
	context.state.(*ConcreteStateA).SwitchStateB(context) // 切换到具体状态B
	context.Handle()

	context.state.(*ConcreteStateB).SwitchStateA(context) // 切换回具体状态A
	context.Handle()
}
