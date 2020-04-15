package var_test

import (
	"testing"
)

const (
	Mondy = 1 + iota
	Tuesday
	Thurday
)

func TestFabi(t *testing.T) {
	a := 1
	b := 1
	t.Log(Mondy)
	t.Log(Tuesday)
	t.Log(Thurday)
	for i := 0;i < 6;i++ {
		a, b = b, a + b
		t.Log(b)
	}
}
