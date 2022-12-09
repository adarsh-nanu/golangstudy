package main

import (
	"adsmod1/communication"
	"adsmod1/emotions"
	"fmt"
	"time"
)

func main() {
	fmt.Println("this is from adsmod1")
	fmt.Println(emotions.Happy())
	//go communication.Server()
	//go communication.Client()
	//time.Sleep(1 * time.Second)
	go communication.Writetochannels()
	go communication.Readfromchannels()
	communication.Timeout()
	communication.Concurrent()
	time.Sleep(30 * time.Second)
}
