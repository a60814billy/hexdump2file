package main

import (
	"fmt"
	"github.com/a60814billy/hexdump2file/internal"
	"github.com/docopt/docopt-go"
	"io/ioutil"
	"os"
	"path/filepath"
)

var usage = `Usage: hexdump2file [-h] [-o OUTPUT_FILE_NAME] [INPUT_FILE_NAME]

Options:
	-h: show this usage
	-o: specific output filename, default is [INPUT_FILE_NAME].bin
`

func ErrorHandler() {
	if err := recover(); err != nil {
		fmt.Printf("Error:\n\t%v\n", err)
		os.Exit(1)
	}
}

func isFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

func main() {
	defer ErrorHandler()
	args, _ := docopt.ParseDoc(usage)

	inputFile, err := args.String("INPUT_FILE_NAME")
	outputFile, err := args.String("OUTPUT_FILE_NAME")

	if len(inputFile) <= 0 {
		println(usage)
		os.Exit(1)
	}
	if len(outputFile) <= 0 {
		filename := filepath.Base(inputFile)
		ext := filepath.Ext(filename)
		filename = filename[0 : len(filename)-len(ext)]
		outputFile = filename + ".bin"
	}

	inputFileFullPath, err := filepath.Abs(inputFile)
	if err != nil {
		panic(err)
	}

	if !isFileExists(inputFileFullPath) {
		panic(fmt.Sprintf("file \"%s\" not exists in %s", inputFile, inputFileFullPath))
	}

	binData := internal.ParseFile(inputFileFullPath)

	if err := ioutil.WriteFile(outputFile, binData, 0644); err != nil {
		panic(err)
	}
}
