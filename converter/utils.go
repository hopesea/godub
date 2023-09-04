package converter

import (
	"fmt"
	"runtime"
	"os/exec"
)

func GetEncoderName() string {
	if IsCommandAvailable(FFMPEGEncoder) {
		return FFMPEGEncoder
	} else {
		panic(fmt.Sprintf("command `%s` not found", FFMPEGEncoder))
	}
}

func IsCommandAvailable(name string) bool {
	checkExistCmd := "which"
	if runtime.GOOS == "windows" {
		checkExistCmd = "where"
	}
	cmd := exec.Command(checkExistCmd, name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
