package txtax

type TransactionTaxInfo struct {
	Transaction Transaction `json:"transaction"`
	CGL         float32     `json:"cgl"`  // capital gain / loss
	Left        float32     `json:"left"` // how much tokens left after calculations
	Error       error       `json:"error"`
}
