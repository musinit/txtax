package txtax_test

import (
	"github.com/musinit/txtax"
	"math/rand"
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

// for future tests
