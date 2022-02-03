# concurrent-ping
Simple and fun concurrent ping utility for multiple hosts. 
Uses linux cli utility "ping" underneath to stream command output concurrently in a non-blocking way.

## Run
```
go run main.go 1.1.1.1 1.0.0.1 8.8.8.8
```
CTRL-C to terminate

## install and run

go install github.com/ChandanGhosh/concurrent-ping@latest

```
concurrent-ping  1.1.1.1 1.0.0.1 8.8.8.8
```
CTRL-C to terminate
