package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Ping utility for multiple hosts, IPs etc.")
	pingCh := make(chan string)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	for _, arg := range os.Args[1:] {
		go pingHost(arg, pingCh)
	}

	for {
		select {
		case <-sigCh:
			os.Exit(-1)
		case v := <-pingCh:
			println(v)
		}
	}

}

func pingHost(arg string, pingCh chan string) {
	cmd := exec.Command("ping", arg)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		pingCh <- m
	}

	cmd.Wait()
}
