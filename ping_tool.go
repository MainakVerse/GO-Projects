package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func pingCmd(host string, count string) *exec.Cmd {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("ping", "-n", count, host)
	default: // linux,darwin
		return exec.Command("ping", "-c", count, host)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ping_tool.go <host> [count]")
		return
	}
	host := os.Args[1]
	count := "4"
	if len(os.Args) >= 3 {
		count = os.Args[2]
	}

	cmd := pingCmd(host, count)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Start error:", err)
		return
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		// simple cleanup so output is brief
		fmt.Println(strings.TrimSpace(line))
	}
	if err := cmd.Wait(); err != nil {
		// non-zero exit may still produce useful output
		// print error silently
	}
}
