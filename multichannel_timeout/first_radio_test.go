package multichannel_timeout

import (
	"testing"
	mulChan "github.com/sunnywalden/go_learning/multichannel_timeout"

)

func TestFirstRadio(t *testing.T) {
	MusicChan := make(chan string, 3)
	NewsChan := make(chan string, 3)
	info := mulChan.FistRadio(MusicChan, NewsChan)
	mulChan.Output("The radio's playing " + info, "")
}
