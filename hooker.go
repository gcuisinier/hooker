package main

import (

	//"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	os.Setenv("HOOKER_DEBUG", "true")
	initLog()
	argsWithProg := os.Args

	executable := argsWithProg[0]

	executable = filepath.Base(executable)
	if executable == "hooker" {
		debugf("Handle hookerctl")

	} else {

		debugf("handle command", executable)

		modifyPath(executable)

		foundExec, _ := exec.LookPath(executable)

		findPreExecHook(executable)

		execute(foundExec, argsWithProg[1:]...)

		findPostExecHook(executable)
	}
}

func findPreExecHook(executable string) {

	homeDir := os.Getenv("HOME")

	info, _ := os.Stat(homeDir + "/.hooker/" + executable + ".preExec")

	debug("pre-hook found", homeDir, info)

}

func findPostExecHook(executable string) {

	homeDir := os.Getenv("HOME")

	info, _ := os.Stat(homeDir + "/.hooker/" + executable + ".preExec")

	debug("post-hook found", homeDir, info)

}

func execute(executable string, arguments ...string) {
	cmd := exec.Command(executable, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
	cmd.Wait()
}

func modifyPath(currentExecutable string) {
	wrapperHomePath, _ := filepath.Abs(filepath.Dir(currentExecutable))

	path := os.Getenv("PATH")
	debugf("Path before : %s", path)
	path = strings.Replace(path, wrapperHomePath+":", "", 1)
	os.Setenv("PATH", path)
	debugf("Path after : %s", path)

}
