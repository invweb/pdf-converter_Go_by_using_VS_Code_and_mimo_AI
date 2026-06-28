# PDF to Text Converter

A Go utility for converting PDF files to text files.

## Installation

```bash
git clone https://github.com/invweb/pdf-converter.git
cd pdf-converter
go build -o pdf-converter
```

## Usage

```bash
# Convert PDF to text (creates .txt file next to original)
./pdf-converter input.pdf

# Specify output filename
./pdf-converter input.pdf output.txt
```

## Example

```bash
$ ./pdf-converter test.pdf
Converted: test.pdf -> test.txt

$ cat test.txt
Hello PDF!

This is a test PDF file created for testing the Go PDF converter.
If you can read this text after conversion, everything works correctly!
```

## Dependencies

- [github.com/ledongthuc/pdf](https://github.com/ledongthuc/pdf) — PDF reading
- [github.com/jung-kurt/gofpdf](https://github.com/jung-kurt/gofpdf) — test PDF creation

## License

MIT
