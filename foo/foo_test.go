package foo

import "testing"

func TestFoo(t *testing.T) {
	actual := Foo()
	if actual != "foo" {
		t.Fail()
	}
}
