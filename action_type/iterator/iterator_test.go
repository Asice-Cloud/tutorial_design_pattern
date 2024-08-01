package iterator

import (
	"testing"
)

func Test(t *testing.T) {
	for v := range Generate() {
		t.Log(v)
	}
}
