package cp_util

import (
	"github.com/greensJadeSoup/v5-go-component/cp_log"
	"os/exec"
	"runtime"
	"strings"
)

func RunInLinuxWithErr(cmd string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", nil
	}

	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		cp_log.Error(err.Error())
		return "", err
	}

	return strings.TrimSpace(string(result)), err
}
