package master_topics

import (
	"fmt"
	"log"
	"time"
)

func tickerCounterOriginal(ticker *time.Ticker, done chan bool) {
	i := 0
Loop:
	for {
		select {
		case t := <-ticker.C:
			i++
			fmt.Println("* Count ", i, " at ", t)
		case <-done:
			fmt.Println("done signal")
			break Loop
		}
	}
	log.Print("Exiting Ticker Counter")

}

func StartTimersTickers() {
	log.Print("\n Starting Ticker Counter \n")
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)
	go tickerCounterOriginal(ticker, done)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	//time.Sleep(3 * time.Second)
	log.Print("Exiting Main")
}
