package foo

import "testing"

func TestIsOne(t *testing.T) {
	expect := true
	result := IsOne(1)

	if expect != result {
		t.Fatal("test")
	}
}
