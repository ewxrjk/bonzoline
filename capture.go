package bonzoline

import (
	"bytes"
	"os"
	"os/exec"
)

// Capture the output of a command
func Capture(name string, arg ...string) (output []byte, err error) {
	cmd := exec.Command(name, arg...)
	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return
	}
	output = outputBuffer.Bytes()
	return
}
