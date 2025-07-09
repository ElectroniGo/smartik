package textextractor

import (
	"bytes"
	"image/png"
	"io"
	"net/http"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/otiai10/gosseract/v2"
)

func ReadScript(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func ExtractText(pdfBuf []byte) (string, error) {
	document, err := fitz.NewFromMemory(pdfBuf)
	if err != nil {
		return "", err
	}
	defer document.Close()

	var scriptTextBuf strings.Builder

	for i := 0; i < document.NumPage(); i++ {
		img, err := document.Image(i)
		if err != nil {
			return "", err
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return "", err
		}

		client := gosseract.NewClient()
		client.SetLanguage("eng")
		defer client.Close()

		if err := client.SetImageFromBytes(buf.Bytes()); err != nil {
			return "", err
		}

		text, err := client.Text()
		if err != nil {
			return "", err
		}

		scriptTextBuf.WriteString(text)
	}

	return scriptTextBuf.String(), nil
}
