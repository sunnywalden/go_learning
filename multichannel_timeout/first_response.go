package multichannel_timeout

import (
	mulChan "github.com/sunnywalden/go_learning/multichannel_timeout"
)

//func FistRadio(MusicChan, NewsChan chan string) string {
//	MusicChannel(MusicChan)
//	NewsChannel(NewsChan)
//	select {
//		case music := <- MusicChan:
//			return music
//			case news := <- NewsChan:
//				return news
//	default:
//		return "No channel"
//	}
//}


func FistRadio(MusicChan, NewsChan chan string) string {
	FirstChan := make(chan string)
	mulChan.MusicChannel(MusicChan)
	mulChan.NewsChannel(NewsChan)
	AllChan := []chan string{MusicChan, NewsChan}
	for _,CurrentChan := range AllChan {
		go func(CurrentChan chan string,FirstChan chan string) {
			res := <- CurrentChan
			FirstChan <- res
		}(CurrentChan,FirstChan,)
	}
	return <-FirstChan
}
