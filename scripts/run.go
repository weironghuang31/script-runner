package scripts

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const IsWindows = runtime.GOOS == "windows"

var shellPath string

func (spec *Spec) Run(names []string) error {
	// set envs, so shell can read these values
	if IsWindows {
		os.Setenv("WINDOWS_SHELL", "1")
	} else {
		os.Setenv("WINDOWS_SHELL", "0")
	}

	for key, value := range spec.Envs {
		os.Setenv(key, value)
	}

	for _, name := range names {
		if err := spec.runScript(name); err != nil {
			return err
		}
	}

	fmt.Println("> Done")

	return nil
}

func (spec *Spec) runScript(name string) error {
	script, ok := spec.Scripts[name]

	if !ok {
		return errors.New("No script: " + name)
	}

	shell, err := getShell()
	if err != nil {
		return err
	}

	command := exec.Command(shell)
	command.Dir = spec.Dir
	command.Stdout = spec.Stdout
	command.Stderr = spec.Stderr
	command.Stdin = bytes.NewBufferString(script)

	fmt.Printf("> Run %v\n", name)

	return command.Run()
}

func getShell() (string, error) {
	if shellPath == "" {
		if IsWindows {
			// find git.exe path, and use the path to find sh.exe
			gitPath, err := exec.LookPath("git.exe")
			if err != nil {
				return "", err
			}

			paths := []string{gitPath, "..", ".."}

			if strings.Contains(gitPath, "mingw64") {
				// if the path contains mingw64 go up one parent
				paths = append(paths, "..")
			}

			paths = append(paths, "bin", "bash.exe")

			shellPath = filepath.Join(paths...)
		} else {
			shPath, err := exec.LookPath("bash")

			if err != nil {
				return "", err
			}

			shellPath = shPath
		}
	}

	return shellPath, nil
}
