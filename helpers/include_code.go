package helpers

import (
	"bytes"
	"os/exec"

	jww "github.com/spf13/jwalterweatherman"
)

func IncludeCode(path string) string {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("cat", path)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		jww.ERROR.Print(stderr.String())
		return stderr.String()
	}
	return out.String()
}
