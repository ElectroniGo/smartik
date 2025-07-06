package textextractor

import (
	"bytes"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gen2brain/go-fitz"
)

func readScript(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch PDF: %v", err)
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read PDF: %v", err)
	}

	return buf
}

func extractText(pdfBuf []byte) (string, error) {
	document, err := fitz.NewFromMemory(pdfBuf)
	if err != nil {
		return "", err
	}
	document.Close()

	var _ strings.Builder // Store extracted text

	for i := 0; i < document.NumPage(); i++ {
		img, err := document.Image(i)
		if err != nil {
			return "", err
		}

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, img, nil); err != nil {
			return "", err
		}

		// create tesseract client
		// read text from image buffer
		// return extracted text
	}

	return "", nil
}
