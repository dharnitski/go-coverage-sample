package bar

import "testing"

func TestBar(t *testing.T) {
	actual := Bar()
	if actual != "bar" {
		t.Fail()
	}
}
