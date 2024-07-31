package composite

import "fmt"

/*
组合模式（Composite Pattern）是一种常用的对象结构型设计模式。它是将对象组合成树状结构以表示“部分-整体”的层次结构，使得客户端以一致的方式处理单个对象和对象组合。

二、组合模式的特点和优点
组合模式的主要优点有：

可以清楚地定义分层次的复杂对象，表示对象的全部或部分层次，使得增加新构件也更容易。
提供一个统一的接口，使得访问组合部分和单个对象的接口一致，使客户端对单个对象和组合对象的使用是一致的。
三、组合模式的应用场景
你希望表示对象的部分-整体层次
你希望客户忽略组合对象与单个对象的不同，客户将统一的使用组合结构中的所有对象。
*/

// Component 是组合模式中的组件接口，定义了搜索方法。
type Component interface {
	// Search 在组件中搜索指定的关键字。
	Search(string)
}

// Product 是具体的组件，实现了 Component 接口。
type Product struct {
	Name string
}

// Search 实现了 Component 接口的方法，在产品中搜索指定的关键字。
func (p *Product) Search(s string) {
	println("search product: %s, keyword: %s", p.Name, s)
}

// Category 是组合对象，包含多个子组件，实现了 Component 接口。
type Category struct {
	Name     string
	Children []Component
}

// Add 向 Category 中添加子组件。
// 参数:
// - child: 要添加的子组件
func (c *Category) Add(child Component) {
	c.Children = append(c.Children, child)
}

// Remove 从 Category 中移除指定位置的子组件。
// 参数:
// - i: 要移除的子组件的索引
func (c *Category) Remove(i int) {
	c.Children = append(c.Children[:i], c.Children[i+1:]...)
}

// Search 实现了 Component 接口的方法，在 Category 及其子组件中搜索指定的关键字。
func (c *Category) Search(s string) {
	fmt.Printf("Category: %s\n", c.Name)
	for _, child := range c.Children {
		child.Search(s)
	}
}
