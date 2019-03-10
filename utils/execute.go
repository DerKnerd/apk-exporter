package utils

import (
	"os/exec"
)

func ExecuteApk(parameters ...string) ([]byte, error) {
	return Execute("apk", parameters...)
}

func Execute(appname string, args ...string) ([]byte, error) {
	command := exec.Command(appname, args...)

	return command.Output()
}
