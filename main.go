package main

import (
	"os"
	"github.com/sparrc/go-ping"
	"net"
	"fmt"
	"strconv"
)

const (
	threadLimit = 5000
)

func main() {
	target, threads := parseArgs()	
	fmt.Println(target, threads)

}

func parseArgs() (target string, numbOfThreads int) {
	if len(os.Args) != 3 { 
		fmt.Println("Error: Ensure you pass two arguments. Host/IP and the number of threads to spawn.")
		os.Exit(1)
	}

	target = os.Args[1]
	numbOfThreads, _ = strconv.Atoi(os.Args[2])

	if ( isHostValid(target) ) == false {
		fmt.Println(" Error: The host you entered is invalid.")
		os.Exit(1)
	}

	if ( areNumberOfThreadsValid(numbOfThreads) ) == false {
		fmt.Println("Error: The number of threads you enter exceeds the limit; ", threadLimit)
		os.Exit(1)
	}

	return target, numbOfThreads 
}

func isHostValid(host string) bool {
	_ , err := net.LookupHost(host)
	if err != nil { return false }

	return true
}

func areNumberOfThreadsValid(threads int) bool {
	threads = int(threads)

	if threads > threadLimit {
		return false
	}
	return true
}
