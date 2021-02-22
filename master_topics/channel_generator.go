package master_topics

import (
	"fmt"
	"log"
	"time"
)

/*
	GO CONCURRENCY REVIEW
	* Composition for multiple independent operations
	* Goroutine: An independently scheduled function, cheaper than regular OS threads
	* Channels: typed way for goroutines communication
	* Channels blocks/lock/sync till data is sent
	* Buffered (unlike unbuffered) channels block only when the buffer is full

	CHANNEL GENERATOR DESIGN PATTERN
	* Function or method that return channels
	* Removes the need to define the channel outside
	* Enables more efficient and fluid code
*/


// now its returning a channel
func tickerCounterChannelGenerator(ticker *time.Ticker) chan bool {
	done := make(chan bool)
	go func() {
		i := 0
	Loop:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("* Count ", i, " at ", t)
			case <-done:
				fmt.Println("done signal")
				break Loop
			}
		}
		log.Print("Exiting Ticker Counter")
	}()
	return done
}

func StartChannelGenerator() {
	log.Print("\n Starting ChannelGenerator \n")
	ticker := time.NewTicker(1 * time.Second)
	// waits for the returning channel
	done := tickerCounterChannelGenerator(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	//time.Sleep(3 * time.Second)
	log.Print("Exiting Main")
}
