package main

import (
	"fmt"
	"github.com/jrroman/de-cli/fs"
)

func main() {
	fs := new(fs.FS)
	contents, err := fs.LsDir("/Users/john/tmp")
	if err != nil {
		fmt.Println(err)
	}

	for _, dir := range contents {
		fmt.Println(dir)
	}
}