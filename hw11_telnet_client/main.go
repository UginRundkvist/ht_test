package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	address := flag.String("timeout", "localhost:4242", "timeout")
	//address := "localhost:4242"
	//timeout := 10 * time.Second
	flag.Parse()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	client := NewTelnetClient(*address, *timeout, os.Stdin, os.Stdout)
	fmt.Println("Подключено")
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
