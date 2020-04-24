package multichannel_timeout

import (
	"testing"
)

func TestAllRadio(t *testing.T) {
	MusicChan := make(chan string, 5)
	NewsChan := make(chan string, 2)
	AllChan := make(chan string, 5)

	MusicChannel(MusicChan)
	NewsChannel(NewsChan)

	MusicToChan(MusicChan, AllChan)
	NewsToChan(NewsChan, AllChan)
	ListenRadio(AllChan)

	//info, err := ListenRadio(AllChan)
	//if err != "" {
	//	} else {
	//		Output(info, "")
	//		}
	//		close(AllChan)
}
