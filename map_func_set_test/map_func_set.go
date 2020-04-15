package map_func_set_test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[string]func(pam int) int{}
	m1["itself"] = func(pam int) int{return pam}
	m1["multi"] = func(pam int) int{return pam * pam}
	nums := []int{1, 2, 3}
	for v := range nums {
		t.Log(m1["itself"](v))
	}
}
