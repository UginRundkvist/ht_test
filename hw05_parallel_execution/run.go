package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return fmt.Errorf("number of gorutins must be greater than zero")
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	var wg sync.WaitGroup
	taskChan := make(chan func() error)
	stopChan := make(chan struct{})
	errChan := make(chan error, len(tasks))
	flg := 0

	// создание горутин
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskChan {
				err := task()
				if err != nil {
					errChan <- err
				}
			}
		}()
	}

	// запись задач
	go func() {
		defer close(taskChan)
		for _, task := range tasks {
			select {
			case taskChan <- task:
			case <-stopChan:
				return
			}
		}
	}()

	// ошибки
	go func() {
		errcount := 0
		defer close(stopChan)
		for err := range errChan {
			if err != nil {
				errcount++
				if errcount >= m {
					flg++
					return
				}
			}
		}
	}()

	wg.Wait()
	close(errChan)
	<-stopChan
	if flg == 1 {
		fmt.Println("ДА")
		return ErrErrorsLimitExceeded
	}

	return nil
}
