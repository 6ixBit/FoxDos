# FoxDos 
HTTP(s) attacks using tcp. [Research purposes only]

## Run - go run main.go <threads> <target> <port>
  Examples:
 - go run main.go 400 www.mysite.com 80
 - go run main.go 400 www.example.com 443

[Warning]: Executing too many threads will crash your computer if it fails to handle the load. On a server/good pc you can feel free to increase the thread count to the thousands as it will increase the amount of requests per second.

## Build
You could also build the executable to binary for cross platform execution by running the commands below.
 - [Windows]: GOOS=windows go build 
 - [Linux]: GOARCH=386 go build

