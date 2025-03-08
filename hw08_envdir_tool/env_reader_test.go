package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "envtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	os.WriteFile(filepath.Join(dir, "FOO"), []byte("bar\n"), 0644)
	os.WriteFile(filepath.Join(dir, "EMPTY"), []byte(""), 0644)

	env, err := ReadDir(dir)
	if err != nil {
		t.Fatalf("Error reading directory: %v", err)
	}

	if env["FOO"].Value != "bar" {
		t.Errorf("Expected 'bar', got '%s'", env["FOO"].Value)
	}
	if !env["EMPTY"].NeedRemove {
		t.Errorf("Expected EMPTY to be marked for removal")
	}
}
