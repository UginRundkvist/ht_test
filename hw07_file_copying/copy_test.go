package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

func TestCopy(t *testing.T) {
	tofile := "testdata/tofile.txt"
	fromf := "testdata/input.txt"

	t.Run("First", func(t *testing.T) {
		testfile1 := "testdata/out_offset0_limit0.txt"
		Copy(fromf, tofile, 0, 0)
		fromfile, _ := os.Stat(testfile1)
		toside, _ := os.Stat(tofile)
		require.Equal(t, toside.Size(), fromfile.Size())
	})

	t.Run("Second", func(t *testing.T) {
		testfile1 := "testdata/out_offset0_limit10.txt"
		Copy(fromf, tofile, 0, 10)
		fromfile, _ := os.Stat(testfile1)
		toside, _ := os.Stat(tofile)
		require.Equal(t, toside.Size(), fromfile.Size())
	})

	t.Run("Third", func(t *testing.T) {
		testfile1 := "testdata/out_offset0_limit1000.txt"
		Copy(fromf, tofile, 0, 1000)
		fromfile, _ := os.Stat(testfile1)
		toside, _ := os.Stat(tofile)
		require.Equal(t, toside.Size(), fromfile.Size())
	})

	t.Run("Fourth", func(t *testing.T) {
		testfile1 := "testdata/out_offset0_limit10000.txt"
		Copy(fromf, tofile, 0, 10000)
		fromfile, _ := os.Stat(testfile1)
		toside, _ := os.Stat(tofile)
		require.Equal(t, toside.Size(), fromfile.Size())
	})

	t.Run("Fifth", func(t *testing.T) {
		testfile1 := "testdata/out_offset100_limit1000.txt"
		Copy(fromf, tofile, 100, 1000)
		fromfile, _ := os.Stat(testfile1)
		toside, _ := os.Stat(tofile)
		require.Equal(t, toside.Size(), fromfile.Size())
	})

	t.Run("Six", func(t *testing.T) {
		testfile1 := "testdata/out_offset6000_limit1000.txt"
		Copy(fromf, tofile, 6000, 10000)
		fromfile, _ := os.Stat(testfile1)
		toside, _ := os.Stat(tofile)
		require.Equal(t, toside.Size(), fromfile.Size())
	})
}
