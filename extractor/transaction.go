package extractor

type Transaction struct {
	PostingDate   string
	EffectiveDate string
	Description   string
	Debit         float64
	Credit        float64
	Balance       float64
}

type Summary struct {
	TotalDebit    float64
	TotalCredit   float64
	EndingBalance float64
}
