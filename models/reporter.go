package models

type HTMLParserInterface interface{
	CreateHtml(templateName string, data interface{}) (string, error)
}

type PDFGeneratorInterface interface{
	CreatePDF(htmlFile string) (string, error)
}

type TemplatePDF struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
}