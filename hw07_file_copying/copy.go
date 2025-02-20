package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	stat, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	sizeInBytes := stat.Size()
	if sizeInBytes < offset {
		return errors.New("невалидная ситуация")
	}

	infile, err := os.Open(fromPath)
	if err != nil {
		return errors.New("Не может открыть входящий файл")
	}

	defer infile.Close()

	_, err = infile.Seek(offset, io.SeekStart)
	if err != nil {
		return errors.New("Не может установить указатель на отступ")
	}

	sizeToCopy := limit
	if limit == -1 {
		sizeToCopy = sizeInBytes - offset
	} else if sizeInBytes-offset < limit {
		sizeToCopy = sizeInBytes - offset
	}

	limitedReader := io.LimitReader(infile, limit)

	tofile, err := os.Create(toPath)
	if err != nil {
		return errors.New("Не может создать файл")
	}
	defer tofile.Close()

	var (
		bytesCopied int64
		mutex       sync.Mutex
		done        = make(chan struct{})
	)

	go func() {
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				fmt.Printf("\rКопирование завершено: 100.00%%      \n")
				return
			case <-ticker.C:
				mutex.Lock()
				if sizeToCopy > 0 {
					progress := float64(bytesCopied) / float64(sizeToCopy) * 100
					fmt.Printf("\rПрогресс: %.2f%%", progress)
				} else {
					fmt.Print("\rПрогресс: Неизвестно (размер файла 0) ")
				}
				mutex.Unlock()
			}
		}
	}()

	buffer := make([]byte, 32*1024)
	for {
		n, err := limitedReader.Read(buffer)
		if err != nil && err != io.EOF {
			close(done)
			return err
		}
		if n == 0 {
			break

			written, err := tofile.Write(buffer[:n])
			if err != nil {
				close(done)
				return err
			}

			mutex.Lock()
			bytesCopied += int64(written)
			mutex.Unlock()
		}
		close(done)

	}
	return nil
}
