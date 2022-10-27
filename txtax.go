package txtax

import (
	"sort"
)

func CalculateCGL(transactions []Transaction, taxMethod TaxMethod) (float32, error) {
	sort.Stable(ByTimestamp(transactions))
	var deposits = map[Currency][]Transaction{}
	accumulatedCGL := float32(0)

	for _, transaction := range transactions {
		if transaction.IsDisabled {
			continue
		}
		switch transaction.Type {
		case TransactionTypeGift:
		case TransactionTypeDonation:
		case TransactionTypeStolen:
			continue
		case TransactionTypeReward:
			if transaction.Category == TxCategoryWithdraw {
				accumulatedCGL += transaction.Total()
			}
		case TransactionTypeFork:
		case TransactionTypeAirdrop:
		case TransactionTypePayment:
			if transaction.Category == TxCategoryDeposit {
				tr := transaction
				deposits[tr.Currency] = append(deposits[tr.Currency], tr)
			} else {
				currencyDeposits := deposits[transaction.Currency]
				if len(currencyDeposits) == 0 {
					return -1, ErrDepositTransactionWasNotFound
				}

				var slideDepositTransaction *Transaction
				switch taxMethod {
				case TaxMethodFIFO:
					slideDepositTransaction = &currencyDeposits[0]
				case TaxMethodLIFO:
					slideDepositTransaction = &currencyDeposits[len(currencyDeposits)-1]
				case TaxMethodHIFO:
					slideDepositTransaction = &currencyDeposits[GetHIFOTransactionIDx(currencyDeposits)]
				default:
					slideDepositTransaction = &currencyDeposits[0]
				}
				cgl := float32(0)
				currTransactionAmount := transaction.Amount
				for currTransactionAmount > 0 {
					availableAmount := currTransactionAmount
					if availableAmount > slideDepositTransaction.Amount {
						availableAmount = slideDepositTransaction.Amount
					}
					cgl += transaction.MarketValue*availableAmount - slideDepositTransaction.MarketValue*availableAmount
					slideDepositTransaction.Amount -= availableAmount

					if len(currencyDeposits) > 1 {
						switch taxMethod {
						case TaxMethodFIFO:
							currencyDeposits = currencyDeposits[1:]
							slideDepositTransaction = &currencyDeposits[0]
						case TaxMethodLIFO:
							currencyDeposits = currencyDeposits[:len(currencyDeposits)-1]
							slideDepositTransaction = &currencyDeposits[len(currencyDeposits)-1]
						case TaxMethodHIFO:
							slideDepositTransaction = &currencyDeposits[GetHIFOTransactionIDx(currencyDeposits)]
						}
					}
					currTransactionAmount -= availableAmount

					if currTransactionAmount > 0 && len(currencyDeposits) == 0 {
						return -1, ErrDepositTransactionWasNotFound
					}
				}
				accumulatedCGL += cgl
			}
		}
	}
	return accumulatedCGL, nil
}

func AnalyseCGL(transactions []Transaction, taxMethod TaxMethod) ([]TransactionTaxInfo, error) {
	sort.Stable(ByTimestamp(transactions))
	var deposits = map[Currency][]Transaction{}
	accumulatedCGL := float32(0)
	txTaxInfo := make([]TransactionTaxInfo, len(transactions))

	for i, transaction := range transactions {
		if transaction.IsDisabled {
			continue
		}
		switch transaction.Type {
		case TransactionTypeGift:
		case TransactionTypeDonation:
		case TransactionTypeStolen:
			continue
		case TransactionTypeReward:
			if transaction.Category == TxCategoryWithdraw {
				accumulatedCGL += transaction.Total()
				txTaxInfo[i] = TransactionTaxInfo{
					Transaction: transaction,
					CGL:         transaction.Total(),
				}
			}

		case TransactionTypeFork:
		case TransactionTypeAirdrop:
		case TransactionTypePayment:
			if transaction.Category == TxCategoryDeposit {
				tr := transaction
				deposits[tr.Currency] = append(deposits[tr.Currency], tr)
				txTaxInfo[i] = TransactionTaxInfo{
					Transaction: transaction,
					CGL:         0,
				}
			} else {
				currencyDeposits := deposits[transaction.Currency]
				if len(currencyDeposits) == 0 {
					txTaxInfo[i] = TransactionTaxInfo{
						Transaction: transaction,
						CGL:         0,
						Error:       ErrNoCorrespondingDepositFound,
					}
					continue
				}

				var slideDepositTransaction *Transaction
				switch taxMethod {
				case TaxMethodFIFO:
					slideDepositTransaction = &currencyDeposits[0]
				case TaxMethodLIFO:
					slideDepositTransaction = &currencyDeposits[len(currencyDeposits)-1]
				case TaxMethodHIFO:
					slideDepositTransaction = &currencyDeposits[GetHIFOTransactionIDx(currencyDeposits)]
				default:
					slideDepositTransaction = &currencyDeposits[0]
				}
				cgl := float32(0)
				currTransactionAmount := transaction.Amount
				for currTransactionAmount > 0 {
					availableAmount := currTransactionAmount
					if availableAmount > slideDepositTransaction.Amount {
						availableAmount = slideDepositTransaction.Amount
					}
					cgl += transaction.MarketValue*availableAmount - slideDepositTransaction.MarketValue*availableAmount
					slideDepositTransaction.Amount -= availableAmount

					if len(currencyDeposits) > 1 {
						switch taxMethod {
						case TaxMethodFIFO:
							currencyDeposits = currencyDeposits[1:]
							slideDepositTransaction = &currencyDeposits[0]
						case TaxMethodLIFO:
							currencyDeposits = currencyDeposits[:len(currencyDeposits)-1]
							slideDepositTransaction = &currencyDeposits[len(currencyDeposits)-1]
						case TaxMethodHIFO:
							slideDepositTransaction = &currencyDeposits[GetHIFOTransactionIDx(currencyDeposits)]
						}
					}
					currTransactionAmount -= availableAmount

					if currTransactionAmount > 0 && len(currencyDeposits) == 0 || (currTransactionAmount != 0 && len(currencyDeposits) == 1 && currencyDeposits[0].Amount == 0) {
						txTaxInfo[i] = TransactionTaxInfo{
							Transaction: transaction,
							CGL:         0,
							Error:       ErrNoCorrespondingDepositFound,
						}
						break
					}
				}

				accumulatedCGL += cgl
				if txTaxInfo[i].Error == nil {
					txTaxInfo[i] = TransactionTaxInfo{
						Transaction: transaction,
						CGL:         cgl,
					}
				}

			}
		}
	}
	return txTaxInfo, nil
}
