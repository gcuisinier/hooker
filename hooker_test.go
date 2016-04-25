package main

import (
	"os"
	"testing"
)

func TestModifyPath(t *testing.T) {

	backupPath := os.Getenv("PATH")
	defer os.Setenv("PATH", backupPath)

	newPath := "/tmp/hooker_test/:" + backupPath

	os.Setenv("PATH", newPath)

	modifyPath("/tmp/hooker_test/hooker")

	reducedPath := os.Getenv("PATH")

	if reducedPath != backupPath {
		t.Errorf("Reduced path %s is not equals to %s", reducedPath, backupPath)
	}

}
