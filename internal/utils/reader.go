package utils

import (
	"bufio"
	"io"
	"os"
)

type Reader struct {
	inputFile    *os.File
	readerObject io.Reader
}

func NewReader(inputFileName string) (reader *Reader, err error) {
	if inputFileName == "" {
		return &Reader{inputFile: nil, readerObject: os.Stdin}, nil
	}
	var inputFile *os.File
	inputFile, err = os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	return &Reader{inputFile: inputFile, readerObject: inputFile}, nil
}

func (reader Reader) ReadLines() (lines []string, err error) {
	scanner := bufio.NewScanner(reader.readerObject)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func (reader Reader) Close() {
	if reader.inputFile == nil {
		reader.inputFile.Close()
	}
}
