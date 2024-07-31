package adapter

import "testing"

func Test(t *testing.T) {
	adapter := NewAdapter()
	res := adapter.Request()
	if res != "specific request" {
		t.Fatalf("expect specific request, but got %s", res)
	}
}
