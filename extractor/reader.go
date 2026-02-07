package extractor

import (
	"bytes"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ReadPDF(path string) ([]string, error) {

	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buf bytes.Buffer

	b, err := r.GetPlainText()
	if err != nil {
		return nil, err
	}

	buf.ReadFrom(b)

	// normalize line endings
	text := strings.ReplaceAll(buf.String(), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	lines := strings.Split(text, "\n")

	var cleaned []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}

	return cleaned, nil
}
