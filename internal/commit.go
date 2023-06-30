package internal

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func Commit(verbose bool) error {
	files := flag.Args()[2:]
	for _, f := range files {
		// check whether file exists
		if err := checkFileExists(f); err != nil {
			return err
		}
	}
	return nil
}

func checkFileExists(path string) error {
	if _, err := os.Stat(path); err == nil {
	} else if os.IsNotExist(err) {
		msg := fmt.Sprintf("File %s does not exist.", path)
		return errors.New(msg)
	} else if os.IsPermission(err) {
		msg := fmt.Sprintf("Permission denied for file %s.", path)
		return errors.New(msg)
	}
	return nil
}
