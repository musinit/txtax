package txtax_test

import (
	"github.com/musinit/txtax"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	txn1 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      2,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      1,
			MarketValue: 200,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
)

func Test_Analyser_Txn1_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn1, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txs1))
}
