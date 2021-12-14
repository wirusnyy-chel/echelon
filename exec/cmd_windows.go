//go:build windows

package exec

import (
	"context"
	"log"
	"os/exec"
)

func getCMD(ctx context.Context, commands ...string) *exec.Cmd {
	err := exec.Command("cmd", "/c", "chcp", "65001").Run()
	if err != nil {
		log.Println(err)
	}
	cmd := exec.CommandContext(ctx, "CMD", "/C")
	cmd.Args = append(cmd.Args, commands...)
	return cmd
}
