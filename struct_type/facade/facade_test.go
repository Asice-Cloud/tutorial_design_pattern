package facade

import "testing"

func Test(t *testing.T) {
	facade := NewFacade()
	facade.Operation()
}
