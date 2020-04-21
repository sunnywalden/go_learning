package multichannel_timeout

import (
	"fmt"
	"time"
)

func MultiChannel(music string, news string) (string, string) {
	Chan_music,_ := MusicChannel(music)
	Chan_news,_ := NewsChannel(news)
	select {
		case res := <- Chan_music:
			fmt.Printf("You're listening to music channel.\nNext song's %s!", res)
			return res,""
			case res := <- Chan_news:
				fmt.Printf("You're listening to news channel.\n %s!", res)
				return res,""
				case <- time.After(time.Second * 3):
					fmt.Print("Timeout after 3 seconds!")
					res := "Timeout"
					return "",res
	default:
		fmt.Print("No channel available!")
		res := "No channel"
		return "",res
	}
}


func MusicChannel(m string) (chan string,string) {
	MusicChan := make(chan string, 3)

	fmt.Printf("Now you're hearing %s !'\n'", m)
	time.Sleep(time.Second * 5)
	MusicChan <- m

	return MusicChan,""
}

func NewsChannel(n string) (chan string,string) {
	NewsChan := make(chan string, 3)

	fmt.Printf("Head news! %s !'\n'", n)
	time.Sleep(time.Second * 1)
	NewsChan <- n

	return NewsChan,""
}