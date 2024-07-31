package builder

import "testing"

func Test(t *testing.T) {
	builder := NewConcreteBuilder()
	director := &Director{builder: builder}
	director.Construct()
	product := builder.getResult()
	product.Show()
}
