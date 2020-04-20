package coroutine_test_test

import (
	"go_learning/coroutine_test"
	"testing"
)

func TestCoroutine(t *testing.T) {
	allNames := []string{"Henry Zhang", "David Beckham", "Cristiano Ronaldo", "Mesut Ozil"}
	coroutine_test.PrintPlayer(allNames)
}
