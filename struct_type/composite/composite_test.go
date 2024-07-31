package composite

import "testing"

func Test(t *testing.T) {
	root := &Category{Name: "Root"}
	electronics := &Category{Name: "Electronics"}

	phone := &Product{Name: "Phone"}
	computer := &Product{Name: "Computer"}

	root.Add(electronics)
	electronics.Add(phone)
	electronics.Add(computer)

	root.Search("phone")
	root.Search("computer")
}
