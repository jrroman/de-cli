package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

type FS struct {}

func (fs *FS) CreateDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("Error creating directory: %s\nError: %s\n", path, err)
	}
	return nil
}

func (fs *FS) LsDir(path string) (dirs []os.FileInfo, err error) {
	return ioutil.ReadDir(path)
}

func main() {
	fs := new(FS)
	dirs, err := fs.LsDir("/home/john/")
	if err != nil {
		fmt.Println(err)
	}

	for _, dir := range dirs {
		fmt.Println(dir)
	}
}
