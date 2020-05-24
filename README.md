# FoxDos
HTTP GET Flood attack with an additional ICMP option

## Run 
  - go run main.go www.mysite.com 100

[Warning]: Executing too many threads will crash your computer if it fails to handle the load

## Build
You could also build the executable to binary for cross platform execution by running the commands below.
 - [Windows]: GOOS=windows go build 
 - [Linux]: GOARCH=386 go build

