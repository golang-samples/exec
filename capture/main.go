package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("ls", "-la").CombinedOutput()
	fmt.Println(string(out))
}
