package main

import (
	"adsmod1/communication"
	"adsmod1/emotions"
	"fmt"
)

func main() {
	fmt.Println("this is from adsmod1")
	fmt.Println(emotions.Happy())
	go communication.Sever()
	//go communication.Client()
}
