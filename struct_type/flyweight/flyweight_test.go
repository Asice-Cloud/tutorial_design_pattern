package flyweight

import "testing"

func Test(t *testing.T) {
	factory := &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}

	flyweight1 := factory.GetFlyweight("A")
	flyweight1.Operation("外部状态1")

	flyweight2 := factory.GetFlyweight("B")
	flyweight2.Operation("外部状态2")

	client := &Client{
		flyweight: factory.GetFlyweight("A"),
	}
	client.Operation()
}
