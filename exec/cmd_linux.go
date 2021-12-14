//go:build linux

package exec

import (
	"context"
	"os/exec"
)

func getCMD(ctx context.Context, commands ...string) *exec.Cmd {
	return exec.CommandContext(ctx, commands...)
}
