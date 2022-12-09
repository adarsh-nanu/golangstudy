package communication

import (
	"fmt"
	"time"
)

//var waitgroup = sync.WaitGroup

var inputchannel = make(chan int)

func Writetochannels() {
	time.Sleep(1 * time.Second)
	fmt.Println("Write")
	inputchannel <- 1
	fmt.Println("Write")
	inputchannel <- 2
	fmt.Println("Write")
	inputchannel <- 3
	fmt.Println("Write")
	inputchannel <- 4
	time.Sleep(2 * time.Second)
	close(inputchannel)
	fmt.Println("end of Writetochannels")
}

func Readfromchannels() {
	for data := range inputchannel {
		fmt.Println("Read ", data)
	}
	fmt.Println("end of Readfromchannels")
}

func Timeout() {
	eventch := make(chan int)
	timeoutch := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		timeoutch <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		eventch <- 1
		eventch <- 2
	}()

	select {
	case <-timeoutch:
		fmt.Println("timedout encountered")
		break
	case data := <-eventch:
		fmt.Println("event happened ", data)
	}
	fmt.Println("end of Timeout")
}

func Concurrent() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan bool)

	go func() {
		for i := 0; i < 15; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			ch2 <- "hi"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			ch3 <- true
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case data := <-ch1:
				fmt.Println("ch1 ", data)
				break
			case data := <-ch2:
				fmt.Println("ch2 ", data)
				break
			case data := <-ch3:
				fmt.Println("ch3 ", data)
				break
			}
			fmt.Println("end of Concurrent")
		}
	}()
}
