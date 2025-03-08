package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	env := make(Environment)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(dir, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		var value string
		if scanner.Scan() {
			value = scanner.Text()
		}
		value = strings.TrimRight(value, " \t")
		value = strings.ReplaceAll(value, "\x00", "\n")

		env[file.Name()] = EnvValue{
			Value:      value,
			NeedRemove: len(value) == 0,
		}
	}

	return env, nil
}
