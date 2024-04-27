package main

import (
	"flag"
	filehandler "huffman/fileHandler"
	"log"
)

func main() {
	var inputFileName string
	var outputFileName string
	flag.StringVar(&inputFileName, "i", "input/input.txt", "input file path")
	flag.StringVar(&outputFileName, "o", "input/test.txt", "output file path")
	flag.Parse()

	_, content := filehandler.ReadFile(inputFileName)
	log.Println(string(content))
}
