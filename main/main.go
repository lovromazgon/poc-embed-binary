// This file runs all binaries in "plugins" folder.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("hello main!")

	plugins, err := os.Open("plugins")
	if err != nil {
		panic(err)
	}

	files, err := plugins.ReadDir(0)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		cmd := exec.Command("./plugins/" + f.Name())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("bye!")
}
