package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeoutPtr := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()
	//localhostPtr := flag.Int("localhost", 4242, "localhost")
	//fmt.Println(*address)

	timeout := *timeoutPtr
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf("usage: go-telnet --timeout=10s host port")
	}
	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	//address := "localhost:4242"
	//timeout := 10 * time.Second
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		fmt.Println("Ошибка подключения к серверу ")
		//fmt.Fprintf(os.Stderr, "Ошибка подключения к серверу  %v", err)
		return
	}

	done := make(chan struct{})
	go func() {
		if err := client.Send(); err != nil {
			fmt.Println("Ошибка при отправке")
		}
		done <- struct{}{}
	}()

	go func() {
		if err := client.Receive(); err != nil {
			if err != io.EOF {
				fmt.Println("Ошибка при получении")
			}
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("Соединение закрыто")
	case <-signalChan:
		fmt.Println("Завершение работы")
	}

	client.Close()
}
