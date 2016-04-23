package main

import (

	//"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("hooker")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func main() {

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Formatter)

	argsWithProg := os.Args

	executable := argsWithProg[0]

	executable = filepath.Base(executable)
	if executable == "hooker" {
		log.Debugf("Handle hookerctl")

	} else {

		log.Debugf("handle command", executable)

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

	log.Debug("pre-hook found", homeDir, info)

}

func findPostExecHook(executable string) {

	homeDir := os.Getenv("HOME")

	info, _ := os.Stat(homeDir + "/.hooker/" + executable + ".preExec")

	log.Debug("post-hook found", homeDir, info)

}

func execute(executable string, arguments ...string) {
	cmd := exec.Command(executable, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
	cmd.Wait()
}

func modifyPath(currentExecutable string) {
	wrapperHomePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	path := os.Getenv("PATH")
	log.Debugf("Path before : %s", path)
	path = strings.Replace(path, wrapperHomePath+":", "", 1)
	os.Setenv("PATH", path)
	log.Debugf("Path after : %s", path)

}
