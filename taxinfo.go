package txtax

type TransactionTaxInfo struct {
	Transaction Transaction `json:"transaction"`
	CGL         float32     `json:"cgl"`
	Error       error       `json:"error"`
}
