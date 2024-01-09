package utils

import (
	"os/exec"
	"strings"
)

func BashExecutor(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ""
	}

	retv := string(out)
	retv = strings.TrimSuffix(retv, "\n")
	return retv
}
