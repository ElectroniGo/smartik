# `core`

This is the service that is responsible for processing PDFs, images & text extraction.

## Prerequisites

This guide assumes you have completed the setup described in [`docs/get-started.md`](../../docs/get-started.md) to prepare your development environment.

## Dependencies

### Tesseract OCR Installation

This service requires Tesseract OCR for text extraction from images and scanned PDFs.

#### Ubuntu/Debian (Linux mint as well)
```bash
# Update package list
sudo apt update
```

```bash
# Install Tesseract and development libraries
sudo apt install tesseract-ocr libtesseract-dev libleptonica-dev tesseract-ocr-eng
```

```bash
# Verify installation
tesseract --version
tesseract --list-langs
```

#### Fedora
```bash
sudo dnf update
```

```bash
# Install Tesseract and development libraries
sudo dnf install tesseract tesseract-devel tesseract-langpack-eng
```

```bash
# Verify installation
tesseract --version
tesseract --list-langs
```

### Go Dependencies

> Make sure you run all the go commands while within the root of the
> core service in [`server/core/`](../../server/core/)

Install the required Go packages:

> Running `pnpm install` as indicated by `get-started.md` will run this command for you but you are welcome to run it again just in case.

```bash
# Install PDF processing library
go mod tidy 
```

## Setup

1. **Configure environment:**
   - Copy `.env.example` to `.env` (if available)
   - Update configuration values as needed

2. **Run the service:**
   ```bash
   go run main.go
   ```

## Configuration

The service requires a PDF URL to be configured. Check the `config` package for available configuration options.

## Troubleshooting

### Tesseract Not Found
Ensure Tesseract is properly installed and accessible:
```bash
which tesseract
tesseract --version
```

If not found, reinstall following the installation steps above.

### No Text Extracted
If OCR returns empty results:
1. Verify the PDF contains scanned images (not native text)
2. Check image quality and resolution
3. Try different language packs if text is not in English
4. Add debug logging to identify the issue

## API Endpoints

*TODO: Document API endpoints when web server is enabled*

## Testing

*Comes after hackathon*