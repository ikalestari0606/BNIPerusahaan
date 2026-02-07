package extractor

import "strings"

func ParseBNI(lines []string) ([]Transaction, Summary) {

	var txns []Transaction
	var cur *Transaction
	var summary Summary
	var stage, sumStage string

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		// ================= SUMMARY LABEL =================
		if line == "Ending Balance :" {
			sumStage = "ending"
			continue
		}
		if line == "Total Debet :" {
			sumStage = "debit"
			continue
		}
		if line == "Total Credit :" {
			sumStage = "credit"
			continue
		}

		// ================= SUMMARY VALUE =================
		if sumStage != "" {
			if isMoney(line) && len(line) > 3 {
				val := parseMoney(line)
				switch sumStage {
				case "ending":
					summary.EndingBalance = val
				case "debit":
					summary.TotalDebit = val
				case "credit":
					summary.TotalCredit = val
				}
				sumStage = ""
			}
			continue
		}

		// ================= TRANSAKSI BARU =================
		if isDateTime(line) && stage == "" {
			if cur != nil {
				txns = append(txns, *cur)
			}
			cur = &Transaction{
				PostingDate: line[:10],
			}
			stage = "effective"
			continue
		}

		if cur == nil {
			continue
		}

		// ================= STATE MACHINE =================
		switch stage {

		case "effective":
			if isDateTime(line) {
				cur.EffectiveDate = line[:10]
				stage = "desc"
			}

		case "desc":
			if isMoney(line) {
				cur.Debit = parseMoney(line)
				stage = "credit"
			} else {
				cur.Description += " " + line
			}

		case "credit":
			if isMoney(line) {
				cur.Credit = parseMoney(line)
				stage = "dk"
			}

		case "dk":
			if line == "D" {
				cur.Credit = 0
				stage = "balance"
			} else if line == "K" {
				cur.Debit = 0
				stage = "balance"
			}

		case "balance":
			// HANYA isi balance jika format SALDO ASLI
			if isBalance(line) && cur.Balance == 0 {
				cur.Balance = parseMoney(line)
			}
			// setelah balance (atau tidak ada), transaksi SELESAI
			stage = ""
		}
	}

	if cur != nil {
		txns = append(txns, *cur)
	}

	return txns, summary
}
