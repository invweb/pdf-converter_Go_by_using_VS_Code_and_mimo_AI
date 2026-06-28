package main

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 24)
	pdf.Cell(40, 10, "Hello PDF!")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 14)
	pdf.MultiCell(0, 7, "This is a test PDF file created for testing the Go PDF converter.\n\nIf you can read this text after conversion, everything works correctly!", "", "", false)

	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		fmt.Printf("Error creating PDF: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created: test.pdf")
}
