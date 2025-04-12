package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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
		return ErrOffsetExceedsFileSize
	}

	infile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer infile.Close()

	_, err = infile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	var sizeToCopy int64
	if limit > 0 && limit <= sizeInBytes-offset {
		sizeToCopy = limit
	} else {
		sizeToCopy = sizeInBytes - offset
	}

	tofile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer tofile.Close()

	buffer := make([]byte, 1024*8) // 8KB буфер
	var copied int64
	start := time.Now()

	for {
		if sizeToCopy > 0 && copied >= sizeToCopy {
			break
		}

		bytesToRead := len(buffer)
		if sizeToCopy > 0 && copied+int64(bytesToRead) > sizeToCopy {
			bytesToRead = int(sizeToCopy - copied)
		}

		n, readErr := infile.Read(buffer[:bytesToRead])
		if n > 0 {
			written, writeErr := tofile.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}
			copied += int64(written)

			printProgress(copied, sizeToCopy)
		}

		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			return readErr
		}
	}

	fmt.Printf("\nКопирование завершено за %v\n", time.Since(start))
	return nil
}

func printProgress(current, total int64) {
	const progressWidth = 40
	percent := float64(current) / float64(total)
	done := int(percent * progressWidth)
	bar := fmt.Sprintf("\r[%s%s] %.2f%% (%d / %d bytes)",
		repeat("=", done),
		repeat(" ", progressWidth-done),
		percent*100,
		current,
		total,
	)
	fmt.Print(bar)
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
