package filehandler

import (
	"io"
	"os"
)

func ReadFile(filename string) (error, []byte) {
	file, err := os.Open(filename)
	if err != nil {
		return err, []byte{}
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err, []byte{}
	}

	return nil, content
}
