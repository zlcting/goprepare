package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		err    error
		output []byte
	)
	cmd = exec.Command("/bin/bash", "-c", "echo 1; echo 2")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))
}
