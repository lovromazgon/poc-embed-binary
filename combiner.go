// This file embeds another binary file, the embedded binary is executed when running main.

package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
)

//go:embed embedded/main
var embeddedFs embed.FS

//go:embed main/main
var mainFs embed.FS

func main() {
	fmt.Println("extracting embedded files")

	f, err := embeddedFs.Open("embedded/main")
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("plugins", os.ModeDir | 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	err = extract(f, "plugins/embedded")
	if err != nil {
		panic(err)
	}

	f, err = mainFs.Open("main/main")
	if err != nil {
		panic(err)
	}

	err = extract(f, "combiner")
	if err != nil {
		panic(err)
	}

	fmt.Println("overwritten combiner, running the new combiner\n")

	cmd := exec.Command("./combiner")
	cmd.Args = os.Args
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			os.Exit(99)
		}
		os.Exit(exitErr.ExitCode())
	}
}

func extract(f fs.File, path string) error {
	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, bytes, 0700)
	if err != nil {
		return err
	}
	return nil
}
