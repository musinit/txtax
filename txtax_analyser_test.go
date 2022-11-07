package txtax_test

import (
	"fmt"
	"github.com/musinit/txtax"
	"github.com/stretchr/testify/assert"
	"math/rand"
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
	txn2 = []txtax.Transaction{
		{
			Hash:        "0",
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "1",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txn3 = []txtax.Transaction{
		{
			Hash:        "0",
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "1",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "2",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 200,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			Hash:        "3",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 300,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			Hash:        "4",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 400,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txn4 = []txtax.Transaction{
		{
			Hash:        "0",
			TimeStamp:   0,
			Amount:      0.1,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "1",
			TimeStamp:   1,
			Amount:      0.1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "2",
			TimeStamp:   2,
			Amount:      0.2,
			MarketValue: 200,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txn5 = []txtax.Transaction{
		{
			Hash:        "0",
			TimeStamp:   0,
			Amount:      0.009,
			MarketValue: 1390,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "1",
			TimeStamp:   1,
			Amount:      0.009,
			MarketValue: 1522,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "2",
			TimeStamp:   2,
			Amount:      0.009,
			MarketValue: 1553,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txn6 = []txtax.Transaction{
		{
			Hash:        "0",
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
			Gas:         1,
		},
		{
			Hash:        "1",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
			Gas:         1,
		},
		{
			Hash:        "2",
			TimeStamp:   2,
			Amount:      1,
			MarketValue: 2000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
			Gas:         1,
		},
	}
)

// 1:10  D
// 1:100 D
// 1:200 W cgl: 190
// 1:300 W сgl: 200
// 1:400 W ?

func Test_Analyser_Txn1_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn1, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txs1))
}

func Test_Analyser_Txn1_HIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn1, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txs1))
}

func Test_Analyser_Txn2_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn2, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn2))
}

func Test_Analyser_Txn2_HIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn2, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn2))
}

func Test_Analyser_Txn3_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn3, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn3))
}

func Test_Analyser_Txn4_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn4, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn4))
}

func Test_Analyser_RandomData_FIFO(t *testing.T) {
	data := make([]txtax.Transaction, 0)
	for i := 0; i < 500; i++ {
		timestamp := int64(i)
		rv := rand.Float32()
		isDeposit := rv > 0.5
		category := txtax.TxCategoryDeposit
		if isDeposit {
			category = txtax.TxCategoryWithdraw
		}
		t := txtax.Transaction{
			TimeStamp:   timestamp,
			Hash:        fmt.Sprintf("hash_%d", i),
			Amount:      rv * 100,
			MarketValue: rv * 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    category,
			Currency:    "ETH",
			IsDisabled:  false,
		}
		data = append(data, t)
	}
	txInfo, err := txtax.AnalyseCGL(data, txtax.TaxMethodFIFO)
	tt := make([]float32, 0)
	for _, t := range txInfo {
		tt = append(tt, t.CGL)
	}

	fmt.Println(tt)
	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(data))
}

func Test_Analyser_Txn4_HIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn4, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn4))
}

func Test_Analyser_Txn5_FIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn5, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn5))
}

func Test_Analyser_Txn5_LIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn5, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn5))
}

func Test_Analyser_RandomData_HIFO(t *testing.T) {
	data := make([]txtax.Transaction, 0)
	for i := 0; i < 500; i++ {
		timestamp := int64(i)
		rv := rand.Float32()
		isDeposit := rv > 0.5
		category := txtax.TxCategoryDeposit
		if isDeposit {
			category = txtax.TxCategoryWithdraw
		}
		t := txtax.Transaction{
			TimeStamp:   timestamp,
			Hash:        fmt.Sprintf("hash_%d", i),
			Amount:      rv * 100,
			MarketValue: rv * 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    category,
			Currency:    "ETH",
			IsDisabled:  false,
		}
		data = append(data, t)
	}
	txInfo, err := txtax.AnalyseCGL(data, txtax.TaxMethodHIFO)
	tt := make([]float32, 0)
	for _, t := range txInfo {
		tt = append(tt, t.CGL)
	}

	fmt.Println(tt)
	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(data))
}

func Test_Analyser_Txn6_LIFO(t *testing.T) {
	txInfo, err := txtax.AnalyseCGL(txn6, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, len(txInfo) == len(txn6))
}
