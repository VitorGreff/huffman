package main

import (
	"flag"
	filehandler "huffman/fileHandler"
	"huffman/frequency"
	"huffman/huff"
	"log"
)

func main() {
	var inputFileName string
	var outputFileName string
	flag.StringVar(&inputFileName, "i", "input/input.txt", "input file path")
	flag.StringVar(&outputFileName, "o", "input/test.txt", "output file path")
	flag.Parse()

	_, content := filehandler.ReadFile(inputFileName)
	m := frequency.BuildFrequencyMap(content)
	tree := huff.BuildTree(m)
	log.Println(tree)
}
