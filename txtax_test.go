package txtax_test

import (
	"github.com/musinit/txtax"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var (
	txs1 = []txtax.Transaction{
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
	txs2 = []txtax.Transaction{
		{
			Hash:        "1",
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			Hash:        "2",
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      1,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   3,
			Amount:      1,
			MarketValue: 2000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs3 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      3,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      1,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   3,
			Amount:      1,
			MarketValue: 2000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs4 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      1,
			MarketValue: 200,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   3,
			Amount:      3,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs5 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      3,
			MarketValue: 10,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
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
	}
	txs6 = []txtax.Transaction{
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
			Amount:      2,
			MarketValue: 50,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "BTC",
		},
		{
			TimeStamp:   2,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   3,
			Amount:      1,
			MarketValue: 500,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "BTC",
		},
		{
			TimeStamp:   4,
			Amount:      1,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "BTC",
		},
		{
			TimeStamp:   5,
			Amount:      1,
			MarketValue: 200,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs7 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "BTC",
		},
		{
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypeGift,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      0.4,
			MarketValue: 10000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   3,
			Amount:      4,
			MarketValue: 250,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   4,
			Amount:      0.9,
			MarketValue: 10000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "BTC",
		},
		{
			TimeStamp:   5,
			Amount:      4,
			MarketValue: 250,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs8 = []txtax.Transaction{
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
			Amount:      2,
			MarketValue: 20,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		}, //10 20 : 200 - 40 = 160 + 100-10=90, 160 + 90 = 250
		{
			TimeStamp:   2,
			Amount:      3,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs9 = []txtax.Transaction{
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
			MarketValue: 20,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      3,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs10 = []txtax.Transaction{
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
			Amount:      2,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      3,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		// FIFO 2000 - 20 = 1980; 1000 - 100 = 900; 1980 + 900 = 2880
		// LIFO 2000 - 200 = 1800; 1000 - 10 = 990; 1800 + 990 = 2790
	}
	txs11 = []txtax.Transaction{
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
			Amount:      3,
			MarketValue: 100,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      3,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
		{
			TimeStamp:   3,
			Amount:      5,
			MarketValue: 2000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   4,
			Amount:      6,
			MarketValue: 1500,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
	txs12 = []txtax.Transaction{
		{
			TimeStamp:   0,
			Amount:      1,
			MarketValue: 1000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "BTC",
		},
		{
			TimeStamp:   1,
			Amount:      1,
			MarketValue: 100,
			Type:        txtax.TransactionTypeGift,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   2,
			Amount:      0.1,
			MarketValue: 10000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "BTC",
		},
		{
			TimeStamp:   4,
			Amount:      4,
			MarketValue: 250,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
		},
		{
			TimeStamp:   5,
			Amount:      4,
			MarketValue: 250,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "ETH",
			IsDisabled:  true,
		},
		{
			TimeStamp:   6,
			Amount:      0.9,
			MarketValue: 100,
			Type:        txtax.TransactionTypeFork,
			Category:    txtax.TxCategoryDeposit,
			Currency:    "BTH",
		},
		{
			TimeStamp:   7,
			Amount:      0.9,
			MarketValue: 10000,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "BTC",
		},
		{
			TimeStamp:   8,
			Amount:      4,
			MarketValue: 250,
			Type:        txtax.TransactionTypePayment,
			Category:    txtax.TxCategoryWithdraw,
			Currency:    "ETH",
		},
	}
)

//| Deposit/Withdraw | Amount | Market Price | Transaction Type | Currency |
//|------------------|--------|--------------|------------------|----------|
//| `Deposit`        | 1      | 1000         | Payment          | BTC      | 1
//| `Deposit`        | 1      | 100          | Gift             | ETH      |
//| `WithDraw`       | 0.1    | 10000        | Payment          | BTC      |1
//| `Deposit`        | 4      | 250          | Payment          | ETH      |1
//| `WithDraw`       | 4      | 250          | Payment          | ETH      |
//| `Deposit`        | 0.9    | 100          | Fork             | BTH      |
//| `WithDraw`       | 0.9    | 10000        | Payment          | BTC      |
//| `WithDraw`       | 4      | 250          | Payment          | ETH      |

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Test_Txs1_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs1, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 280)
}

func Test_Txs2_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs2, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 2890)
}

func Test_Txs3_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs3, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 2980)
}

func Test_Txs4_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs4, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 2690)
}

func Test_Txs5_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs5, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 90)
}

func Test_Txs8_LIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs8, txtax.TaxMethodLIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 250)
}

func Test_Txs9_HIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs9, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 260)
}

func Test_Txs9_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs9, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 260)
}

func Test_Txs10_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs10, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 2880)
}

func Test_Txs10_LIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs10, txtax.TaxMethodLIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 2790)
}

func Test_Txs10_HIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs10, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 2790)
}

func Test_Txs11_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs11, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 3680)
}

func Test_Txs11_LIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs11, txtax.TaxMethodLIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 1690)
}

func Test_Txs11_HIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs11, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 1690)
}

func Test_Txs12_FIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs12, txtax.TaxMethodFIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 9000)
}

func Test_Txs12_LIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs12, txtax.TaxMethodLIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 9000)
}

func Test_Txs12_HIFO(t *testing.T) {
	cg, err := txtax.CalculateCGL(txs12, txtax.TaxMethodHIFO)

	assert.Nil(t, err)
	assert.True(t, cg == 9000)
}
