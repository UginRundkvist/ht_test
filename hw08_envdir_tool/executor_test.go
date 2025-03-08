package main

import "testing"

func TestRunCmd(t *testing.T) {
	env := Environment{
		"FOO": {Value: "bar", NeedRemove: false},
	}

	cmd := []string{"env"}
	returnCode := RunCmd(cmd, env)

	if returnCode != 0 {
		t.Errorf("Expected return code 0, got %d", returnCode)
	}
}
