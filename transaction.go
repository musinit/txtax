package txtax

type Transaction struct {
	TimeStamp   int64           `json:"timeStamp"`
	Hash        string          `json:"hash"`
	Amount      float32         `json:"amount"`
	MarketValue float32         `json:"marketValue"`
	Type        TransactionType `json:"type"`
	Category    TxCategory      `json:"category"`
	Currency    Currency        `json:"currency"`
	IsDisabled  bool            `json:"isActive"`
	Gas         float32         `json:"gas"` // amount of fee in $ that was used for transaction
}

type TransactionType string

var (
	TransactionTypeGift     TransactionType = "GIFT"
	TransactionTypeDonation TransactionType = "DONATION"
	TransactionTypeStolen   TransactionType = "STOLEN"
	TransactionTypeAirdrop  TransactionType = "AIRDROP"
	TransactionTypeFork     TransactionType = "FORK"
	TransactionTypePayment  TransactionType = "PAYMENT"
	TransactionTypeReward   TransactionType = "REWARD"
)

type TxCategory string

var (
	TxCategoryDeposit  TxCategory = "DEPOSIT"
	TxCategoryWithdraw TxCategory = "WITHDRAW"
)

type TaxMethod string

var (
	TaxMethodFIFO TaxMethod = "FIFO"
	TaxMethodLIFO TaxMethod = "LIFO"
	TaxMethodHIFO TaxMethod = "HIFO"
)

type Currency string

func (t *Transaction) Total() float32 {
	return t.Amount*t.MarketValue - t.Gas
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
