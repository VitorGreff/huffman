package main

import (
	"flag"
	filehandler "huffman/fileHandler"
	"huffman/huff"
	"log"
)

func main() {
	var inputFileName string
	var outputFileName string
	flag.StringVar(&inputFileName, "i", "input/input.txt", "input file path")
	flag.StringVar(&outputFileName, "o", "output/test.txt", "output file path")
	flag.Parse()

	// _, content := filehandler.ReadFile(inputFileName)
	// m := frequency.BuildFrequencyMap(content)
	m := map[string]int{
		"C": 32,
		"D": 42,
		"E": 120,
		"K": 7,
		"L": 42,
		"M": 24,
		"U": 37,
		"Z": 2,
	}
	t := huff.BuildTree(m)
	huff.PrintTree(t)
	p := huff.BuildPrefixTable(t, "", map[string]string{})
	log.Println(p)
	filehandler.WriteHeader(p, outputFileName)
	log.Println(filehandler.ReadHeader(outputFileName))
	log.Println(huff.DecodePrefixCode(t, "111100"))
}
