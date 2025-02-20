package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tof := "/home/Ugin/all/otus/ht_test/hw07_file_copying/tofile.txt"
	fromf := "/home/Ugin/all/otus/ht_test/hw07_file_copying/testdata/input.txt"

	// testfile3 := "/home/Ugin/all/otus/ht_test/hw07_file_copying/testdata/out_offset0_limit1000.txt"
	err := Copy(fromf, tof, 0, 1000)

	// toopen, _ := os.ReadFile(fromf)
	// openfile3, _ := os.ReadFile(testfile3)
	// require.True(t, bytes.Equal(toopen, openfile3), "Файлы %s и %s должны быть идентичны", tof, testfile3)
	require.Equal(t, nil, err)
}
