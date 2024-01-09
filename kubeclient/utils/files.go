package utils

import (
	"os"
	"fmt"
)

func WriteFile(path, contents string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(contents))
	if err != nil {
		fmt.Println(err)
	}
	return err
}