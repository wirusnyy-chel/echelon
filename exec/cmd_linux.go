//go:build linux

package exec

import (
	"context"
	"os/exec"
)

func getCMD(ctx context.Context, commads ...string) *exec.Cmd {
	return exec.CommandContext(ctx, commands...)
}
