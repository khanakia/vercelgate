package utils

import (
	"errors"
	"fmt"
	"os"
)

func OpenFile(filepath string) ([]byte, error) {
	info, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) || info.IsDir() {
		fmt.Println("file does not exist")
		return nil, err
	}

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("unable to read file")
		return nil, err
	}

	return fileBytes, nil
}

func IsFileExists(filepath string) error {
	info, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) || info.IsDir() {
		// fmt.Println("file does not exist")
		return err
	}

	_, err = os.ReadFile(filepath)
	if err != nil {
		// fmt.Println("unable to read file")
		return err
	}

	return nil
}
