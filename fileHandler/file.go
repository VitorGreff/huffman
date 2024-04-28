package filehandler

import (
	"encoding/gob"
	"io"
	"os"
)

func ReadFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return []byte{}
	}

	return content
}

func WriteHeader(prefixTable map[string]string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(prefixTable)
	if err != nil {
		panic(err)
	}
}

func ReadHeader(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	var p map[string]string
	err = dec.Decode(&p)
	if err != nil {
		panic(err)
	}

	return p
}
