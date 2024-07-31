package factory_mode

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	factoryA := &ConcreteFactory_A{}
	// 调用工厂方法创建产品对象
	productA := factoryA.CreateProduct()
	fmt.Println(productA.Operation())

	// 创建具体工厂B
	factoryB := &ConcreteFactory_B{}
	// 调用工厂方法创建产品对象
	productB := factoryB.CreateProduct()
	fmt.Println(productB.Operation())
}
