package iterator

/*
迭代器模式是一种行为设计模式，它提供一种统一的方法来遍历一个聚合对象中的各个元素，而无需暴露该聚合对象的内部表示。

2. 迭代器模式的特点和优点
迭代器模式的特点和优点如下：

可以隐藏集合对象的内部结构，使得遍历算法和集合对象解耦。
提供了一种统一的方式来遍历不同类型的集合对象。
简化了客户端代码，使得客户端代码更加清晰、简洁。
可以提供迭代器的不同实现方式，以适应不同的遍历需求。
3. 迭代器模式的实际应用场景举例
迭代器模式在实际应用中有很多场景，例如：

遍历一个数据库查询结果集。
遍历一个文件系统中的文件和文件夹。
遍历一个集合中的元素。

*/

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type ConcreteIterator struct {
	collection *ConcreteCollection
	index      int
}

func (c *ConcreteIterator) HasNext() bool {
	if c.index < len(c.collection.Items) {
		return true
	}
	return false
}

func (c *ConcreteIterator) Next() interface{} {
	if c.HasNext() {
		item := c.collection.Items[c.index]
		c.index++
		return item
	}
	return nil
}

type Collection interface {
	CreateIterator() Iterator
}

type ConcreteCollection struct {
	Items []interface{}
}

func (c *ConcreteCollection) CreateIterator() Iterator {
	return &ConcreteIterator{
		collection: c,
		index:      0,
	}
}

// Generate simply to use iterator
// Generate 生成一个迭代器(yield)
/*
在以上代码中，我们定义了一个 GenerateItems() 函数，它返回一个只读通道 (<-chan)，在这个函数中使用 yield 将元素依次发送到通道中。
使用 range 遍历这个只读通道，并输出元素值。

这样，我们就使用生成器函数简化了迭代器的实现。
*/
func Generate() <-chan interface{} {
	collection := []interface{}{"c++", "java", "python", "golang"}
	out := make(chan interface{}) // if not buffered, you need to use goroutine to send value to channel
	// if without goroutine, you should change "out" into buffered channel: out := make(chan interface{},SIZE)
	go func() {
		for _, v := range collection {
			out <- v
		}
		close(out)
	}()
	return out
}
