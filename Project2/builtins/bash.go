package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Bash(scriptName string) error {
	if scriptName == "" {
		return fmt.Errorf("bash: missing script argument")
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Construct the full path to the script
	scriptPath := filepath.Join(currentDir, scriptName)

	// Check if the script is executable
	scriptInfo, err := os.Stat(scriptPath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if scriptInfo.Mode()&0111 == 0 {
		return fmt.Errorf("bash: script is not executable")
	}

	// Modify the PATH to include the current directory
	currentPath := os.Getenv("PATH")
	currentPath = currentDir + string(filepath.ListSeparator) + currentPath
	os.Setenv("PATH", currentPath)

	// Use Bash to execute the script
	cmd := exec.Command("bash", scriptName)
	// Set the command's working directory to the current directory
	cmd.Dir = currentDir

	// Capture and print the output
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(stdout))

	return nil
}
