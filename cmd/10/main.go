package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout duration")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet --timeout=<timeout> <host> <port>")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", host+":"+port)

	stop := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		close(stop)
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
			if n > 0 {
				_, err = conn.Write(buf[:n])
				if err != nil {
					fmt.Println("Error writing to server:", err)
					return
				}
			}
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Connection closed by server:", err)
			return
		}
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}
