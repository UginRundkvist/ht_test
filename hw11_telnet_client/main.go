package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	timeoutPtr := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()

	timeout := *timeoutPtr
	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf("usage: go-telnet --timeout=10s host port")
	}
	host := args[0]
	port := args[1]
	address := net.JoinHostPort(host, port)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		fmt.Println("Ошибка подключения к серверу ", err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	done := make(chan struct{})
	go func() {
		defer wg.Done()
		if err := client.Send(); err != nil {
			fmt.Println("Ошибка при отправке", err)
		}
		done <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		if err := client.Receive(); err != nil {
			if err != io.EOF {
				fmt.Println("Ошибка при получении ", err)
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
	wg.Wait()
	client.Close()
}
