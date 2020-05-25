# FoxDos 
HTTP GET Flood attack with an additional ICMP option. [Research purposes only]

## Run - <method> <host> <threads>
  Examples:
          - go run main.go httpget www.mysite.com 500
          - go run main.go icmp 127.0.0.1 500

[Warning]: Executing too many threads will crash your computer if it fails to handle the load

## Build
You could also build the executable to binary for cross platform execution by running the commands below.
 - [Windows]: GOOS=windows go build 
 - [Linux]: GOARCH=386 go build

