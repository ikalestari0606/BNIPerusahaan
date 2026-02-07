package exporter

import (
	"bni/extractor"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExportExcel(txns []extractor.Transaction, sum extractor.Summary, path string) error {

	f := excelize.NewFile()
	sheet := "Sheet1"

	// ===== STYLE =====
	numberStyle, _ := f.NewStyle(&excelize.Style{
		NumFmt: 3, // format angka ribuan
	})

	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	descStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
			Vertical: "top",
		},
	})

	// ===== HEADER =====
	headers := []string{"Posting Date", "Effective Date", "Description", "Debit", "Credit", "Balance"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Style header
	f.SetCellStyle(sheet, "A1", "F1", headerStyle)

	// ===== SET LEBAR KOLOM (INI KUNCI KERAPIAN) =====
	f.SetColWidth(sheet, "A", "A", 15) // Posting Date
	f.SetColWidth(sheet, "B", "B", 15) // Effective Date
	f.SetColWidth(sheet, "C", "C", 50) // Description
	f.SetColWidth(sheet, "D", "D", 18) // Debit
	f.SetColWidth(sheet, "E", "E", 18) // Credit
	f.SetColWidth(sheet, "F", "F", 20) // Balance

	// ===== ISI DATA =====
	for r, t := range txns {
		row := r + 2

		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), t.PostingDate)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), t.EffectiveDate)

		descCell := fmt.Sprintf("C%d", row)
		f.SetCellValue(sheet, descCell, t.Description)
		f.SetCellStyle(sheet, descCell, descCell, descStyle)

		if t.Debit != 0 {
			c := fmt.Sprintf("D%d", row)
			f.SetCellValue(sheet, c, t.Debit)
			f.SetCellStyle(sheet, c, c, numberStyle)
		}

		if t.Credit != 0 {
			c := fmt.Sprintf("E%d", row)
			f.SetCellValue(sheet, c, t.Credit)
			f.SetCellStyle(sheet, c, c, numberStyle)
		}

		c := fmt.Sprintf("F%d", row)
		f.SetCellValue(sheet, c, t.Balance)
		f.SetCellStyle(sheet, c, c, numberStyle)
	}

	// Freeze header (biar scroll enak)
	f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		Split:       true,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	return f.SaveAs(path)
}
