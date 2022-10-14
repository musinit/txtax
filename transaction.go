package txtax

type Transaction struct {
	TimeStamp   int             `json:"timeStamp"`
	Hash        string          `json:"hash"`
	Amount      float32         `json:"amount"`
	MarketValue float32         `json:"marketValue"`
	Type        TransactionType `json:"type"`
	Category    TxCategory      `json:"category"`
	Currency    Currency        `json:"currency"`
	IsDisabled  bool            `json:"isActive"`
}

type TransactionType string

var (
	TransactionTypeGift     TransactionType = "gift"
	TransactionTypeDonation TransactionType = "donation"
	TransactionTypeStolen   TransactionType = "stolen"
	TransactionTypeAirdrop  TransactionType = "airdrop"
	TransactionTypeFork     TransactionType = "fork"
	TransactionTypePayment  TransactionType = "payment"
	TransactionTypeReward   TransactionType = "reward"
)

type TxCategory string

var (
	TxCategoryDeposit  TxCategory = "deposit"
	TxCategoryWithdraw TxCategory = "withdraw"
)

type TaxMethod string

var (
	TaxMethodFIFO TaxMethod = "fifo"
	TaxMethodLIFO TaxMethod = "lifo"
	TaxMethodHIFO TaxMethod = "hifo"
)

type Currency string

func (t *Transaction) Total() float32 {
	return t.Amount * t.MarketValue
}

func GetHIFOTransactionIDx(transactions []Transaction) int {
	highest := &transactions[0]
	idx := 0
	for i, value := range transactions {
		if highest.MarketValue < value.MarketValue && value.Amount > 0 {
			highest = &value
			idx = i
		}
	}
	return idx
}
