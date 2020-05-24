package main

import (
	"os"
	 // "github.com/sparrc/go-ping"
	"log"
	"net"
	"fmt"
)

func main() {
	parseArgs()	


}

func parseArgs() (target, numbOfThreads string) {
	if len(os.Args) != 3 { 
		fmt.Println("Error: Ensure you pass two arguments. Host/IP and the number of threads to spawn.")
		os.Exit(1)
	}

	target = os.Args[1]
	numbOfThreads = os.Args[2]

	if ( isHostValid(target) ) == false {
		fmt.Println(" Error: The host you entered is invalid.")
		os.Exit(1)
	}
	return target, numbOfThreads 
}

func isHostValid(host string) bool {
	result, err := net.LookupHost(host)
	if err != nil { return false }

	return true
}