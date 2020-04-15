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

func TestSet(t *testing.T) {
	s1 := map[string]bool{}
	s1["SunnyWalden"] = true
	s1["Henry"] = true
	user1 := "Henry"
	user2 := "Walden"
	if s1[user1] {
		t.Log("User %s already exists", user1)
	} else {
		s1[user2] = true
		t.Log("User %s registry success!", user2)
	}
}
