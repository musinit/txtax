package txtax_test

import (
	"github.com/musinit/txtax"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	txna1 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      2,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 200,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
)

func Test_Analyse_Txn1_FIFO(t *testing.T) {
	incomeTypes := map[txtax.TransactionType]struct{}{
		txtax.TransactionTypePayment: struct{}{},
	}
	cglTypes := map[txtax.TransactionType]struct{}{
		txtax.TransactionTypePayment: struct{}{},
	}
	txInfo, err := txtax.Analyse(txna1, txtax.TaxOptions{
		TaxMethod:   txtax.TaxMethodFIFO,
		IncomeTypes: incomeTypes,
		CGLTypes:    cglTypes,
		SkipTypes:   nil,
	})

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txna1))
}
