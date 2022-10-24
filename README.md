<br>
<img src="https://www.linkpicture.com/q/txtax-logo.png" alt="alt text" height="50">

TxTax is a crypto transactions tax calculator. It analyzes transactions and provides data for tax reports. <br>Let's make smth like CoinTracker or Turbotax, but free and opensource!

### Features
 - Support any type of crypto currencies
 - FIFO, LIFO, HIFO tax generation algorithms
 - Several types of crypto transactions, including:
   - gifts
   - donations
   - stolen/broken
   - airdrops
   - payments
   - forks
 - Manual enabling/disabling transactions from calculating
# Docs

## [TxTax](#)

Crypto transactions taxes calculator

[![Go Test](https://github.com/musinit/txtax/actions/workflows/main.yml/badge.svg)](https://github.com/musinit/txtax/actions/workflows/main.yml)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/musinit/txtax/blob/master/LICENSE)

### Download

```shell
go get -u github.com/musinit/txtax
```

### Getting started
function CalculateCGL is the main function for calculating capital gain/loss.<br>
It espects 2 arguments: 
- the list of transactions (it would be sorting inside by Timestamp)
- Tax calculation method (FIFO, LIFO, HIFO)

Transaction has a simplified structure in order to calculate capital gain/loss.
````
{
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
````


## Test examples
| Category | Amount | Market Price | Transaction Type |
|------------------|--------|--------------|------------------|
| `Deposit`        | 2      | 10           | Payment          |
| `Deposit`        | 2      | 100          | Payment          |
| `WithDraw`       | 3      | 1000         | Payment          |

Taxable profits
 - FIFO<br>
   (2000 - 20) + (1000 - 100) = **2880**
 - LIFO<br>
   (2000 - 200) + (1000 - 10) = **2790**
 - HIFO<br>
   (2000 - 200) + (1000 - 10) = **2790**

<br>

|Category | Amount | Market Price | Transaction Type |
|------------------|--------|--------------|------------------|
| `Deposit`        | 2      | 10           | Payment          |
| `Deposit`        | 3      | 100          | Payment          |
| `WithDraw`       | 3      | 1000         | Payment          |
| `Deposit`        | 5      | 2000         | Payment          |
| `WithDraw`       | 6      | 1500         | Payment          |

Taxable profits
- FIFO<br>
  (2000 - 20) + (1000-100) + (3000-200) + (6000-8000) = **3680**
- LIFO<br>
  (3000 - 300) + (7500-10000) + (1500-10)=**1690**
- HIFO<br>
  the same as LIFO

<br>

Great [Coinbase example](https://www.cointracker.io/blog/crypto-tax-guide#how-is-cryptocurrency-taxed)

| Category   | Amount | Market Price | Transaction Type | Currency |
|------------|--------|--------------|------------------|----------|
| `Deposit`  | 1      | 1000         | Payment          | BTC      |
| `Deposit`  | 1      | 100          | Gift             | ETH      |
| `WithDraw` | 0.1    | 10000        | Payment          | BTC      |
| `Deposit`  | 4      | 250          | Payment          | ETH      |
| `WithDraw` | 4      | 250          | Payment          | ETH      |
| `Deposit`  | 0.9    | 100          | Fork             | BTH      |
| `WithDraw` | 0.9    | 10000        | Payment          | BTC      |
| `WithDraw` | 4      | 250          | Payment          | ETH      |

Taxable profits
- FIFO **9000**
- LIFO **9000**
- HIFO **9000**
