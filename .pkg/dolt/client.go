package dolt

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecutablePath() string {
	path, err := filepath.Abs(EXECUTABLE)
	if err != nil {
		panic(err)
	}
	return path
}

func New() (Interface, error) {
	return nil, nil
}

func Init() {
	fmt.Println(os.Executable())
	cmd := exec.Command(ExecutablePath(), "version")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
