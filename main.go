package main

import (
	"flag"
	"fmt"

	filehandler "huffman/fileHandler"
	"huffman/frequency"
	"huffman/huff"
)

func main() {
	var inputFileName string
	var outputFileName string
	flag.StringVar(&inputFileName, "i", "input/input.txt", "input file path")
	flag.StringVar(&outputFileName, "o", "output/output.txt", "output file path")
	flag.Parse()

	content := filehandler.ReadFile(inputFileName)
	m := frequency.BuildFrequencyMap(content)
	/* m := map[string]int{
		"C": 32,
		"D": 42,
		"E": 120,
		"K": 7,
		"L": 42,
		"M": 24,
		"U": 37,
		"Z": 2,
	} */
	t := huff.BuildTree(m)
	p := huff.BuildPrefixTable(t, "", map[string]string{})

	// ENCODE TO STORE
	encodedContent := huff.EncodeText(t, content)
	filehandler.WriteHeader(p, outputFileName)
	filehandler.WriteEncodedString(outputFileName, encodedContent)

	// DECODE TO LOG
	header, filePrefixCode := filehandler.ReadEncodedFile(outputFileName)
	fmt.Printf("\nPrefix Table -> %v\n", header)
	fmt.Printf("\nEncoded Content -> %s\n", filePrefixCode)
	fmt.Printf("\nDecoded Content -> %s\n", huff.DecodePrefixCode(t, filePrefixCode))
}
