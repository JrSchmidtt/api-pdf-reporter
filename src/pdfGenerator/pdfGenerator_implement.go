package pdfGenerator

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wk struct {
	rootPath string
}

func NewWkHtmlToPdf(rootPath string) PDFGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {
	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	f, err := os.Open(htmlFile)
	if err != nil {
		return "failure open", err
	}

	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "failure new pdf", err
	}

	pdf.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdf.Create(); err != nil {
		return "failure buffer page", err
	}

	if err := pdf.WriteFile(fileName); err != nil {
		return "failure WriteFile", err
	}

	return fileName, nil
}