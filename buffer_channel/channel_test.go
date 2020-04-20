package buffer_channel

import "testing"

func TestChannel(t *testing.T) {
	AllPlayers := []string{"Henry Zhang", "David Beckham", "Cristiano Ronaldo", "Mesut Ozil"}
	AsyncTasks(AllPlayers)
	PrepareMatch()
}
