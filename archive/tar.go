package archive

import (
	"os"
	"fmt"
	"net/http"
	// "archive/tar"
	// "compress/gzip"
)

type Compressed struct {}

func (c *Compressed) DetectCompression(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Error opening file: %s\nError: %s", path, err)
	}
	defer file.Close()

	var buf []byte = make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil {
		return "", fmt.Errorf("Error reading file: %s\nError: %s", file, err)
	}

	cmptype := http.DetectContentType(buf)
	return cmptype, nil
}