package algs

import (
	"fmt"
	"testing"
)

func TestListFiles(t *testing.T) {
	dir := "."

	files, err := ListFiles(dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(files)
}
