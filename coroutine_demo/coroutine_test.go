package coroutine_demo_test

import (
	"github.com/sunnywalden/go_learning/coroutine_demo"
	"testing"
)

func TestCoroutine(t *testing.T) {
	allNames := []string{"Henry Zhang", "David Beckham", "Cristiano Ronaldo", "Mesut Ozil"}
	coroutine_demo.PrintPlayer(allNames)
}
