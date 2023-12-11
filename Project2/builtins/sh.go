package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Shell(name string) error {
	if name == "" {
		return fmt.Errorf("sh: missing argument")
	}

	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	scriptPath := filepath.Join(currentDir, name)

	scriptInfo, err := os.Stat(scriptPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if scriptInfo.Mode()&0111 == 0 {
		return fmt.Errorf("sh: script is not executable")
	}

	currentPath := os.Getenv("PATH")
	currentPath = currentDir + string(filepath.ListSeparator) + currentPath
	os.Setenv("PATH", currentPath)

	cmd := exec.Command("sh", name)
	cmd.Dir = currentDir

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(stdout))

	return nil
}
