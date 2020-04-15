package logic_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
	)

func TestLogic(t *testing.T) {
	a := 7 // 0111
	t.Log(a &^Readable == Readable, a &^Writable == Writable, a &^Executable == Executable)
}
