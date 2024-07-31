package facade

import "fmt"

/*
外观模式（Facade Pattern）是一种结构型设计模式，提供了一个统一的接口，用于访问子系统中的一组接口。它隐藏了子系统的复杂性，为外部提供了一个简化的接口。

2. 外观模式的特点和优点
外观模式有以下特点和优点：

提供了一个简化的接口，使得子系统更易于使用。
将客户端与子系统的耦合度降低，客户端只需要与外观类进行交互，而不需要了解子系统的具体实现细节。
符合开闭原则，可以很方便地增加或修改子系统中的功能。
3. 外观模式的实际应用场景举例
外观模式在实际开发中有着广泛的应用场景，例如：

提供简化的接口来访问复杂的第三方库或API。
封装一组复杂的逻辑操作，以简化客户端的调用过程。
为现有系统提供一个简单的接口，以便与其他系统进行集成
在这个示例中，我们假设有一个电子商务平台，它包括订单管理、库存管理和支付系统
。订单管理系统负责创建订单、查询订单等功能，库存管理系统负责查询商品库存、扣减库存等功能，支付系统负责支付订单的功能。
为了简化客户端与子系统的交互，我们可以使用外观模式来设计这些子系统的接口。
*/

// Facade 是外观类，提供了一个统一的接口来访问子系统。
type Facade struct {
	system_a *SystemA
	system_b *SystemB
	system_c *SystemC
}

// NewFacade 创建并返回一个新的 Facade 实例。
func NewFacade() *Facade {
	return &Facade{
		system_a: NewSystemA(),
		system_b: NewSystemB(),
		system_c: NewSystemC(),
	}
}

// Operation 调用子系统的操作方法。
func (f *Facade) Operation() {
	f.system_a.Operation_A()
	f.system_b.Operation_B()
	f.system_c.Operation_C()
}

// SystemA 是子系统A，包含 Operation_A 方法。
type SystemA struct{}

// NewSystemA 创建并返回一个新的 SystemA 实例。
func NewSystemA() *SystemA {
	return &SystemA{}
}

// Operation_A 是 SystemA 的操作方法。
func (s *SystemA) Operation_A() {
	fmt.Println("Operation_A")
}

// SystemB 是子系统B，包含 Operation_B 方法。
type SystemB struct{}

// NewSystemB 创建并返回一个新的 SystemB 实例。
func NewSystemB() *SystemB {
	return &SystemB{}
}

// Operation_B 是 SystemB 的操作方法。
func (s *SystemB) Operation_B() {
	fmt.Println("Operation_B")
}

// SystemC 是子系统C，包含 Operation_C 方法。
type SystemC struct{}

// NewSystemC 创建并返回一个新的 SystemC 实例。
func NewSystemC() *SystemC {
	return &SystemC{}
}

// Operation_C 是 SystemC 的操作方法。
func (s *SystemC) Operation_C() {
	fmt.Println("Operation_C")
}
