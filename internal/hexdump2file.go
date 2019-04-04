package internal

import (
	"bytes"
	"encoding/hex"
	"os"
	"strings"
	"unsafe"
)

func parseLineData(line string) (bool, []string) {
	if len(line) > 45 && line[8] == ':' {
		return true, strings.Split(line[10:46], " ")
	}
	return false, nil
}

func readAllFileLines(file *os.File) []string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	b := buf.Bytes()
	s := *(*string)(unsafe.Pointer(&b))
	return strings.Split(s, "\n")
}

func ParseFile(filepath string) []byte {
	file, _ := os.Open(filepath)

	lines := readAllFileLines(file)

	buf := new(bytes.Buffer)

	for _, line := range lines {
		if line[len(line)-1:] == " " {
			line = line[:len(line)-2]
		}
		if valid, hexDataArray := parseLineData(line); valid == true {
			for _, data := range hexDataArray {
				rawData, _ := hex.DecodeString(data)
				buf.Write(rawData)
			}
		}
	}

	return buf.Bytes()
}

