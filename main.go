package main

import (
	"fmt"
	// "strings"
	// "github.com/jrroman/de-cli/fs"
	"github.com/jrroman/de-cli/cmd"
)

func main() {
	cmd := new(cmd.RunCmd)
	out, err := cmd.Run("ls", "-a")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}