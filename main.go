package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pdf-converter <input.pdf> [output.txt]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := strings.TrimSuffix(inputFile, ".pdf") + ".txt"
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	f, reader, err := pdf.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening PDF: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var text strings.Builder
	textReader, err := reader.GetPlainText()
	if err != nil {
		fmt.Printf("Error extracting text: %v\n", err)
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	for {
		n, err := textReader.Read(buf)
		if n > 0 {
			text.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}

	err = os.WriteFile(outputFile, []byte(text.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Converted: %s -> %s\n", inputFile, outputFile)
}
