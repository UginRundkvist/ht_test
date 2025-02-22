package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
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

	var sizeToCopy int64
	if limit > 0 {
		sizeToCopy = limit
	} else {
		sizeToCopy = sizeInBytes - offset
	}

	if limit == 0 {
		limit = sizeInBytes
	}

	bar := pb.New64(sizeToCopy)
	bar.SetTemplateString(`{{ etime . }} [{{ bar . }}] {{ percent . }}`)
	bar.Start()

	fmt.Println("Сколько нужно cкопировать байт: ", sizeToCopy)

	limitedReader := bar.NewProxyReader(io.LimitReader(infile, limit))

	tofile, err := os.Create(toPath)
	if err != nil {
		return errors.New("Не может создать файл")
	}
	defer tofile.Close()

	batecopied, err := io.Copy(tofile, limitedReader)
	bar.Finish()
	fmt.Println(batecopied)
	return nil
}
