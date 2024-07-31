package bridge

/*
桥接模式，也称为Bridge模式，是一种结构型设计模式，它将抽象与实现解耦，使得两者可以独立地变化。简单地说，桥接模式是一种让抽象和实现分离的方案。

1.2 桥接模式的目的和作用
桥接模式的主要目的是将抽象部分与其实现部分解耦，以便二者可以独立地进行变化和扩展。这是通过建立一个抽象的桥接类完成的，该类链接到一个具体实现类。

2. 桥接模式的特点和优点
桥接模式的一些主要特点和优点如下：

提高了系统的可扩展性。抽象和实现都可以独立扩展，不会影响到对方。
符合开闭原则，抽象部分和实现部分可以分别独立扩展，而不会影响到彼此。
实现细节对客户透明，可以对用户隐藏实现细节。
3. 桥接模式的应用场景
当你想要分离一个复杂对象的实现和抽象的时候，可以使用桥接模式。 这可能会对现有代码的性能有好影响，尤其是当程序在运行时只使用了对象的一部分。
当你需要在多个对象间分享特定的实现状态，但对于客户方代码而言，它们需要呈现为独立类。
*/

// DrawAPI 是一个抽象接口，定义了绘制圆形的方法。
type DrawAPI interface {
	Draw_A_Circle(radius int, x int, y int)
}

// RedCircle 是 DrawAPI 的具体实现，绘制红色的圆形。
type RedCircle struct{}

// Draw_A_Circle 实现了 DrawAPI 接口的方法，绘制红色的圆形。
func (r *RedCircle) Draw_A_Circle(radius int, x int, y int) {
	println("Drawing Circle[ color: red, radius: ", radius, ", x: ", x, ", y: ", y, "]")
}

// BlueCircle 是 DrawAPI 的具体实现，绘制蓝色的圆形。
type BlueCircle struct{}

// Draw_A_Circle 实现了 DrawAPI 接口的方法，绘制蓝色的圆形。
func (b *BlueCircle) Draw_A_Circle(radius int, x int, y int) {
	println("Drawing Circle[ color: blue, radius: ", radius, ", x: ", x, ", y: ", y, "]")
}

// Shape 是一个接口，定义了绘制形状的方法。
type Shape interface {
	Draw()
}

// Circle 是 Shape 的具体实现，表示一个圆形。
type Circle struct {
	x, y, radius int
	drawAPI      DrawAPI
}

// Draw 实现了 Shape 接口的方法，使用 DrawAPI 绘制圆形。
func (c *Circle) Draw() {
	println("Circle Draw()")
	c.drawAPI.Draw_A_Circle(c.radius, c.x, c.y)
}
