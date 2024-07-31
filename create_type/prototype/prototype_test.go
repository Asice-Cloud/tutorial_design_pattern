package prototype

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	manager := NewPrototypeManager()

	// 在原型管理类中注册原型对象
	manager.Register("productA", &Product{name: "Product A"})

	// 使用原型管理类创建和克隆对象
	productA := manager.Clone("productA").(*Product)
	fmt.Println(productA.name) // Output: Product A
}
