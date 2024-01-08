package utils

import (
	"io"
	"os"
)

type Writer struct {
	outputFile   *os.File
	writerObject io.Writer
}

func NewWriter(outputFileName string) (writer *Writer, err error) {
	if outputFileName == "" {
		return &Writer{outputFile: nil, writerObject: os.Stdout}, nil
	}
	var outputFile *os.File
	outputFile, err = os.Create(outputFileName)
	if err != nil {
		return nil, err
	}
	return &Writer{outputFile: outputFile, writerObject: outputFile}, nil
}

func (writer Writer) WriteLines(lines []string) (err error) {
	for _, line := range lines {
		_, err = writer.writerObject.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (writer Writer) Close() {
	if writer.outputFile != nil {
		writer.outputFile.Close()
	}
}
