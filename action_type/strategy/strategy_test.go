package strategy

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	context := &Context{}

	context.SetStrategy(&ConcreteStrategyA{})
	result := context.ExecuteStrategy("Pay by Alipay") //here strategy is nil
	fmt.Println(result)

	context.SetStrategy(&ConcreteStrategyB{})
	result = context.ExecuteStrategy("Pay by Wechat")
	fmt.Println(result)
}
