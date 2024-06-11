package main

import "gogo/day1-day10/day8/studyclient"

func main() {
	//studygoroutine.FiveGoroutines()
	//studylock.Cal()
	//fmt.Println("Final count:", studylock.Count)
	//studylock.RUnlock()
	//studychannel.Chan()
	//studychannel.CloseChannel()
	//studychannel.RangeChannel()
	//goroutineChannel.CatWriteRead()
	studyclient.CreateServer()
	studyclient.CreateClient()
}
