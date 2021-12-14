//go:build windows

package exec

import (
	"context"
	"log"
	"os/exec"
)

func getCMD(ctx context.Context) *exec.Cmd {
	err := exec.Command("cmd", "/c", "chcp", "65001").Run()
	if err != nil {
		log.Println(err)
	}
	return exec.Command("CMD", "/C")
}
