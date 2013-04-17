package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// start
	cmd := exec.Command("sleep", "5")
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// wait or timeout
	donec := make(chan error, 1)
	go func() {
		donec <- cmd.Wait()
	}()
	select {
	case <-time.After(3 * time.Second):
		cmd.Process.Kill()
		fmt.Println("timeout")
	case <-donec:
		fmt.Println("done")
	}
}
