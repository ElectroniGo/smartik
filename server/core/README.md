# `core`

This is the service that is responsible for processing PDFs, images & text extraction.

> All commands in this file assume, you are running them from the root of this service. (`server/core/`) if not, they won't run properly.

## Prerequisites

This guide assumes you have completed the setup described in [`docs/get-started.md`](../../docs/get-started.md) to prepare your development environment.

### Go Dependencies

Running `pnpm install` as indicated by `get-started.md` will run this command for you but you are welcome to run it again just in case.

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

   The server will be accessible at [http://localhost:1323] if you do not change the default `PORT` value set in [`.env.example`](./env.example)

## Configuration

The service requires a PDF URL to be configured. Check the `config` package for available configuration options.

## Troubleshooting

### Tesseract Not Found
Ensure Tesseract is properly installed and accessible (refer to [docs/getting-started.md](../../docs/get-started.md#2-install-dependencies)):

```bash
which tesseract
tesseract --version
```

### No Text Extracted
If OCR returns empty results:
1. Verify the PDF contains scanned images (not native text)
2. Check image quality and resolution
3. Try different language packs if text is not in English
4. Add debug logging to identify the issue

## API Endpoints

### POST /trigger

**Input:**

```json
{
  "url": "link-to-online-pdf"
}
```

**Response:**

```json
{
  "success": true,
  "detail":  "Text extracted successfully",
  "data":    "text data extracted from pdf",
}
```

or

Response with `success: false` and detail information

## Testing

*Comes after hackathon*