package buffer_channel

import (
	"testing"
	buff "github.com/sunnywalden/go_learning/buffer_channel"
)

func TestChannel(t *testing.T) {
	AllPlayers := []string{"Henry Zhang", "David Beckham", "Cristiano Ronaldo", "Mesut Ozil"}
	buff.AsyncTasks(AllPlayers)
	buff.PrepareMatch()
}
