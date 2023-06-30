package internal

import (
	"fmt"
	"os"
)

// InitRepo initializes a FCV repository.
// It creates a .fcv file in the current working directory.
// TODO use io/fs for testability
func InitRepo(verbose bool) error {
	if verbose {
		fmt.Fprintln(os.Stderr, "initializeing repository")
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = os.Create(cwd + "/.fcv")
	if err != nil {
		return err
	}

	return nil
}
