package exec

import (
	"context"
	"runtime"
	"testing"
)

type test struct {
	os      string
	command string
	stdin   string
	stdout  bool
	stderr  bool
}

var WinTests = []test{
	{os: "windows", command: "dir", stdin: "", stdout: true, stderr: false},
	{os: "windows", command: "abc", stdin: "", stdout: false, stderr: true},
	//{os: "windows", command: "wsl cat", stdin: "123", stdout: true, stderr: false},
	//{os: "windows", command: "wsl cat", stdin: "", stdout: false, stderr: false},
	{os: "linux", command: "dir", stdin: "", stdout: false, stderr: false},
}
var LinuxTests = []test{
	{os: "linux", command: "echo 123", stdin: "", stdout: true, stderr: false},
	{os: "linux", command: "cat", stdin: "112", stdout: true, stderr: false},
	{os: "linux", command: "cat", stdin: "", stdout: false, stderr: false},
	{os: "linux", command: "a", stdin: "", stdout: false, stderr: true},
	{os: "windows", command: "dir", stdin: "", stdout: false, stderr: false},
}

func TestRunCommand(t *testing.T) {
	var tests []test
	if runtime.GOOS == "windows" {
		tests = WinTests
	} else {
		tests = LinuxTests
	}
	for _, v := range tests {
		out, err := RunCommand(context.Background(), v.command, v.os, v.stdin)
		if (out != "") != v.stdout || (err != "") != v.stderr {
			t.Errorf("test:%v,\n result: out:%v,err:%v", v, out, err)
		}
	}

}
