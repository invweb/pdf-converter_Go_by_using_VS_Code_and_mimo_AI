# PDF to Text Converter

Go-утилита для конвертации PDF файлов в текстовые файлы.

## Установка

```bash
git clone https://github.com/invweb/pdf-converter.git
cd pdf-converter
go build -o pdf-converter
```

## Использование

```bash
# Конвертировать PDF в текст (создастся .txt файл рядом)
./pdf-converter input.pdf

# Указать имя выходного файла
./pdf-converter input.pdf output.txt
```

## Пример

```bash
$ ./pdf-converter test.pdf
Converted: test.pdf -> test.txt

$ cat test.txt
Hello PDF!

This is a test PDF file created for testing the Go PDF converter.
If you can read this text after conversion, everything works correctly!
```

## Зависимости

- [github.com/ledongthuc/pdf](https://github.com/ledongthuc/pdf) — чтение PDF
- [github.com/jung-kurt/gofpdf](https://github.com/jung-kurt/gofpdf) — создание тестовых PDF

## Лицензия

MIT
