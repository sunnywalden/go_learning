package multichannel_timeout

import (
	"fmt"
	"time"
)

func MultiChannelListener(MusicChan,NewsChan  chan string) (res string,ending bool) {

	select {
		case res,ok := <- MusicChan:
			if ok {
				fmt.Printf("Next song!\n %s\n", res)
				return res, false
			}
	case <- time.After(time.Second * 5):
		fmt.Print("Timeout after 5 seconds!")
		return "",true
	default:
		fmt.Print("No channel available!")
		//res := "No channel"
		return "",true
	}

	select {

		case res,ok := <- NewsChan:
			if ok {
				fmt.Printf("Next news!.\n %s\n", res)
				return res, false
			}
		case <- time.After(time.Second * 5):
			fmt.Print("Timeout after 5 seconds!")
			return "",true
		default:
			fmt.Print("News channel not available!")
			return "",true
	}

	return "",true
}


func MusicListener(MusicChan chan string) (music string,ending bool) {
	music, ok := <- MusicChan
	if ok {
		fmt.Printf("Now you're hearing %s !'\n'", music)
		time.Sleep(time.Second * 3)
		return music, false
	} else {
		Output("Rest time!", "")
		return "Rest time!", true
	}
}

func NewsListener(NewsChan chan string) (news string,ending bool) {
	news, ok := <- NewsChan
	if ok {
		fmt.Printf("Heading news!\n %s !'\n'", news)
		time.Sleep(time.Second * 3)
		return news, false
	} else {
		Output("Rest time!", "")
		return "Rest time!", true
	}
}