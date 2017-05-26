package main

import (
	"fmt"
	// "github.com/jrroman/de-cli/fs"
	// "github.com/jrroman/de-cli/cmd"
	"github.com/jrroman/de-cli/archive"
)

func main() {
	c := new(archive.Compressed)
	cmp, err := c.CompressionType("")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cmp)
}