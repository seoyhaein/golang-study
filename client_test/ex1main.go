package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	// construct `sleep.sh` command
	cmd := &exec.Cmd{
		Path:   "./sleep.sh",
		Args:   []string{"./sleep.sh", "3"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// run `cmd` in background
	cmd.Start()

	// do something else
	for i := 1; i < 300000; i++ {
		fmt.Println(i)
	}

	// wait `cmd` until it finishes
	cmd.Wait()
}
