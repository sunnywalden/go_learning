package multichannel_timeout

import (
	"testing"

)

func TestFirstRadio(t *testing.T) {
	MusicChan := make(chan string, 3)
	NewsChan := make(chan string, 3)
	info := FistRadio(MusicChan, NewsChan)
	Output("The radio's playing " + info, "")
}
