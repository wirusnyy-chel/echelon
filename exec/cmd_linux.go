//go:build linux

package exec

import (
	"context"
	"os/exec"
)

func getCMD(ctx context.Context) *exec.Cmd {
	return exec.Command("bash")
}
