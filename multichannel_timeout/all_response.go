package multichannel_timeout

import (
	mulChan "github.com/sunnywalden/go_learning/multichannel_timeout"

)

func AllRadio(MusicChan, NewsChan chan string, ResChan chan string) {


	mulChan.MusicChannel(MusicChan)
	mulChan.NewsChannel(NewsChan)

}

func MusicToChan(MusicChan chan string, ResChan chan string)  {
	//result := ""
	for i:=0;i<= len(MusicChan);i++ {
		res,ending := mulChan.MusicListener(MusicChan)
		if ending != "" {
			close(MusicChan)
			break
		} else {
			//result += res
			ResChan <- res
		}
	}
}

func NewsToChan(NewsChan chan string, ResChan chan string)  {
	//result := ""
	for i:=0;i<= len(NewsChan);i++ {
		res,ending := mulChan.NewsListener(NewsChan)
		if ending != "" {
			close(NewsChan)
			break
		} else {
			//result += res
			ResChan <- res
		}
	}
}

func ListenRadio(AllChan chan string)  {

	for i := 0;i< len(AllChan);i++ {
		info, ok := <-AllChan
		if ! ok {
			mulChan.Output("No channel!", "")
			//return "", "No channel"
			//break
		} else {
			mulChan.Output("Radio signal! "+info, "")
			//return info, ""
		}
	}
	//return info, ""
}
