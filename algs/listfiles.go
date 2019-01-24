package algs

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFiles(dir string) ([]string, error) {
	fileList := make([]string, 16)
	e := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil
}
