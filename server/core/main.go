package main

import (
	"github.com/smartik/core/config"
	textextractor "github.com/smartik/core/internal/text_extractor"
)

var cfg = config.LoadConfig()
var pdfURL = cfg.PdfURL

func main() {
	// Uncomment the following lines to set up an Echo server with logging and recovery middleware.
	// This is commented out to focus on the text extraction functionality.
	//
	// e := echo.New()

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// e.Use(middleware.Recover())

	// routes.Attatch(e)

	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))

	pdfData := textextractor.ReadScript(pdfURL)
	println("PDF data size:", len(pdfData), "bytes")

	scriptText, err := textextractor.ExtractText(pdfData)
	if err != nil {
		panic(err)
	}

	println("Extracted text length:", len(scriptText))
	println("Extracted text:")
	println(scriptText)

	if len(scriptText) == 0 {
		println("No text was extracted!")
	}
}
