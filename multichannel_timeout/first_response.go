package multichannel_timeout

func FistRadio(MusicChan, NewsChan chan string) string {
	MusicChannel(MusicChan)
	NewsChannel(NewsChan)
	select {
		case music := <- MusicChan:
			return music
			case news := <- NewsChan:
				return news
	default:
		return "No channel"
	}
}
