package util

import (
	"errors"
	"path/filepath"
)

// Validating input file
func ValidateFile(file string) error {
	if file == "" {
		return errors.New("File name is empty")
	} else {
		ext := filepath.Ext(file)
		if ext != ".json" && ext != ".yaml" && ext != ".yml" {
			return errors.New("Invalid file: " + file)
		}
	}
	return nil
}
