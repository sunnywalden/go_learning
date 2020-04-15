package string_test

import (
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	name := "张博"
	uni_name := []rune(name)
	t.Logf("%s %d %x %d", name, len(name), uni_name, len(uni_name))
	for _,s:= range name {
		t.Logf("%c", s)
	}
}


func TestStrConv(t *testing.T) {
	age := 30
	worked_years := "7"
	t.Logf("My age is %s", strconv.Itoa(age))
	if wy, err := strconv.Atoi(worked_years);err == nil {
		t.Logf("I've worked for %d years", wy)
	} else {
		wy := 0
		t.Logf("Using default configure!")
		t.Logf("I've worked for %d years", wy)
	}
}