package multichannel_timeout

import (
	"testing"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func TestMusicChannel(t *testing.T) {
	MusicChan := make(chan string, 5)
	MusicChannel(MusicChan)
	for {
		res,ending := MusicListener(MusicChan)
		if ending != "" {
			break
		} else {
			Output("Music for now." + res + "\n", "")
		}
	}

}


func TestNewsChannel(t *testing.T) {

	NewsChan := make(chan string, 5)
	NewsChannel(NewsChan)
	for {
		res,ending := NewsListener(NewsChan)
		if ending != "" {
			break
		} else {
			Output("News for now." + res +  "\n", "")
		}
	}

}


//func TestMultiChannel(t *testing.T) {
//	NewsChan := make(chan string, 5)
//	MusicChan := make(chan string, 5)
//	NewsChannel(NewsChan)
//	MusicChannel(MusicChan)
//	for {
//		res,ending := MultiChannelListener(MusicChan,NewsChan)
//		if ending {
//			break
//		} else {
//			fmt.Printf("For now. %s\n", res)
//		}
//	}
//}