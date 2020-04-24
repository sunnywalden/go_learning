package multichannel_timeout

import (
	"fmt"
	"time"
)

func MultiChannelListener(MusicChan,NewsChan  chan string) (res string,err string) {

	select {
		case res,ok := <- MusicChan:
			if ok {
				fmt.Printf("Next song!\n %s\n", res)
				return res, ""
			}
	case <- time.After(time.Second * 5):
		fmt.Print("Timeout after 5 seconds!")
		return "","Timeout"
	default:
		fmt.Print("No channel available!")
		return "", "channel not available"
	}

	select {

		case res,ok := <- NewsChan:
			if ok {
				fmt.Printf("Next news!.\n %s\n", res)
				return res, ""
			}
		case <- time.After(time.Second * 5):
			fmt.Print("Timeout after 5 seconds!")
			return "","Timeout"
		default:
			fmt.Print("News channel not available!")
			return "","channel not available"
	}

	return "","Rest time"
}


func MusicListener(MusicChan chan string) (res string, err string) {
	//go func(MusicChan chan string) (res string, err string) {
			music, ok := <-MusicChan
			if ok {
					Output("Next music! " + music + "!\n", "")
					time.Sleep(time.Second * 3)
					return music, ""
				} else {
					//close(MusicChan)
					Output("Rest time!", "")
					return "", "Rest time!"
					}

	//}(MusicChan,)
	//return "", "Rest time!"
}

func NewsListener(NewsChan chan string) (news string,err string) {
	//go func(MusicChan chan string) (res string, err string) {
		news, ok := <-NewsChan
		if ok {
			Output("Heading news!\n" + news + "!\n", "")
			time.Sleep(time.Second * 3)
			return news, ""
		} else {
			//close(NewsChan)
			Output("Rest time!", "")
			return "", "Rest time!"
		}
	//}(NewsChan)
	//return "", "Rest time!"
}