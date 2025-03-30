package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3" //nolint:depguard
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	stat, err := os.Stat(fromPath)
	if err != nil {
		fmt.Println("Ошибка с stat")
		return err
	}

	sizeInBytes := stat.Size()
	if sizeInBytes < offset {
		fmt.Println("Ошибка с sizeInBytes")
		return err
	}

	infile, err := os.Open(fromPath)
	if err != nil {
		fmt.Println("Ошибка с infile")
		return err
	}

	defer infile.Close()

	_, err = infile.Seek(offset, io.SeekStart)
	if err != nil {
		fmt.Println("Ошибка с Seek")
		return err

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
		fmt.Println("Ошибка с tofile")
		return err
	}
	defer tofile.Close()
	batecopied, err := io.Copy(tofile, limitedReader)
	if err != nil {
		fmt.Println("Ошибка с batecopied")
		return err
	}
	bar.Finish()
	fmt.Println(batecopied)
	return nil
}
