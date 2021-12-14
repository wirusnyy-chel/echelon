package exec

import (
	"bytes"
	"context"
	"log"
	"runtime"
	"strings"
)

func RunCommand(ctx context.Context, command, os, stdin string) (string, string) {
	if os != runtime.GOOS {
		return "", ""
	}
	cmd := getCMD(ctx, strings.Fields(command)...)
	var stdout, stderr bytes.Buffer
	stdinReader := strings.NewReader(stdin)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = stdinReader
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	return stdout.String(), stderr.String()
}
