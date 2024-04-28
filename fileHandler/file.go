package filehandler

import (
	"encoding/gob"
	"io"
	"os"
	"strings"
)

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return ""
	}

	return string(content)
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

	_, err = file.WriteString("\nHEADER DELIMITER")
	if err != nil {
		panic(err)
	}

}

func ReadEncodedFile(filename string) (map[string]string, string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	sliceContent := strings.Split(string(content), "HEADER DELIMITER")

	var p map[string]string
	// reads the header
	dec := gob.NewDecoder(strings.NewReader(sliceContent[0]))
	err = dec.Decode(&p)
	if err != nil {
		panic(err)
	}

	prefixCode := sliceContent[1]

	return p, prefixCode
}

func WriteEncodedString(filename string, prefixCode string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + prefixCode)
	if err != nil {
		panic(err)
	}
}
