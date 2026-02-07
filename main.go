package main

import (
	"bni/exporter"
	"bni/extractor"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// =============================
	// INPUT & OUTPUT DI FOLDER YANG SAMA DENGAN EXE
	// =============================
	inputDir := "input"
	outputDir := "output"

	// Pastikan folder ada
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".pdf") {
			continue
		}

		pdfPath := filepath.Join(inputDir, file.Name())

		lines, err := extractor.ReadPDF(pdfPath)
		if err != nil {
			log.Println("ReadPDF error:", err)
			continue
		}

		txns, sum := extractor.ParseBNI(lines)

		outPath := filepath.Join(
			outputDir,
			strings.TrimSuffix(file.Name(), ".pdf")+".xlsx",
		)

		err = exporter.ExportExcel(txns, sum, outPath)
		if err != nil {
			log.Println("ExportExcel error:", err)
			continue
		}

		log.Println("Sukses:", file.Name())
	}
}
