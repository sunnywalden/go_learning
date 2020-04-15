package condition_test

import "testing"

func TestCondition(t *testing.T) {
		//i := 4
	for i := 0;i < 5;i++ {
		if i < 2 {
			t.Log("Primary ")
		} else if i < 4 {
			t.Log("Middle")

		} else {
			t.Log("High")
		}
	}
}

func TestSwitch(t *testing.T) {
	for i := 0;i < 5;i++ {
		switch i % 2 {
		case 0:
			t.Log("Even")
		case 1:
			t.Log("Odd")
		default:
			t.Log("Others")
		}
	}
}
