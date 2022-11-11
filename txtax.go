package txtax

import (
	"errors"
	"sort"
)

func AnalyseCGL(transactions []Transaction, taxMethod TaxMethod) ([]TransactionTaxInfo, error) {
	sort.Stable(ByTimestamp(transactions))
	var deposits = map[Currency][]Transaction{}
	var supportDepositHash = map[string]int{}
	accumulatedCGL := float32(0)
	txTaxInfo := make([]TransactionTaxInfo, len(transactions))

	for i, transaction := range transactions {
		idx := i
		if transaction.IsDisabled {
			continue
		}
		switch transaction.Type {
		case TransactionTypeStolen:
			continue
		case TransactionTypeDonation:
		case TransactionTypeGift:
		case TransactionTypeReward:
			if transaction.Category == TxCategoryWithdraw {
				accumulatedCGL += transaction.Total()
				txTaxInfo[i] = TransactionTaxInfo{
					Transaction: transaction,
					CGL:         transaction.Total(),
					Left:        transaction.Amount,
				}
			}

		case TransactionTypeFork:
		case TransactionTypeAirdrop:
		case TransactionTypePayment:
			if transaction.Category == TxCategoryDeposit {
				tr := transaction
				deposits[tr.Currency] = append(deposits[tr.Currency], tr)
				supportDepositHash[transaction.Hash] = idx
				txTaxInfo[i] = TransactionTaxInfo{
					Transaction: transaction,
					CGL:         0,
					Left:        transaction.Amount,
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
				idx := len(currencyDeposits)
				for currTransactionAmount > 0 {
					availableAmount := currTransactionAmount
					if availableAmount > slideDepositTransaction.Amount {
						availableAmount = slideDepositTransaction.Amount
					}
					cgl += transaction.MarketValue*availableAmount - slideDepositTransaction.MarketValue*availableAmount - slideDepositTransaction.Gas
					slideDepositTransaction.Amount -= availableAmount

					slideIDs, ok := supportDepositHash[slideDepositTransaction.Hash]
					if !ok {
						return nil, errors.New("can't find corresponding deposit hash in supportDeposits")
					}
					txTaxInfo[slideIDs].Left -= availableAmount

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
					idx--

					if idx < 0 || (currTransactionAmount > 0 && len(currencyDeposits) == 0 || (currTransactionAmount != 0 && len(currencyDeposits) == 1 && currencyDeposits[0].Amount == 0)) {
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
						CGL:         cgl - transaction.Gas,
					}
				}

			}
		}
	}
	return txTaxInfo, nil
}

func Analyse(transactions []Transaction, options TaxOptions) ([]TransactionTaxInfo, error) {
	sort.Stable(ByTimestamp(transactions))
	var deposits = map[Currency][]Transaction{}
	var supportDepositHash = map[string]int{}
	accumulatedCGL := float32(0)
	txTaxInfo := make([]TransactionTaxInfo, len(transactions))

	for i, transaction := range transactions {
		idx := i
		tr := transaction
		if transaction.IsDisabled {
			continue
		}
		// skip
		if _, ok := options.SkipTypes[transaction.Type]; ok {
			continue
		}

		// income
		if _, ok := options.IncomeTypes[transaction.Type]; ok {
			if transaction.Category == TxCategoryDeposit {
				deposits[tr.Currency] = append(deposits[tr.Currency], tr)
				supportDepositHash[transaction.Hash] = idx
				txTaxInfo[i] = TransactionTaxInfo{
					Transaction: transaction,
					CGL:         0,
					Income:      transaction.Amount * transaction.MarketValue,
					Left:        transaction.Amount,
				}
			}
		}

		// cgl
		if _, ok := options.CGLTypes[transaction.Type]; ok {
			if transaction.Category == TxCategoryWithdraw {
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
				switch options.TaxMethod {
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
				idx := len(currencyDeposits)
				for currTransactionAmount > 0 {
					availableAmount := currTransactionAmount
					if availableAmount > slideDepositTransaction.Amount {
						availableAmount = slideDepositTransaction.Amount
					}
					cgl += transaction.MarketValue*availableAmount - slideDepositTransaction.MarketValue*availableAmount - slideDepositTransaction.Gas
					slideDepositTransaction.Amount -= availableAmount

					slideIDs, ok := supportDepositHash[slideDepositTransaction.Hash]
					if !ok {
						return nil, errors.New("can't find corresponding deposit hash in supportDeposits")
					}
					mv := txTaxInfo[slideIDs].Income / txTaxInfo[slideIDs].Left
					txTaxInfo[slideIDs].Left -= availableAmount
					txTaxInfo[slideIDs].Income = txTaxInfo[slideIDs].Left * mv

					if len(currencyDeposits) > 1 {
						switch options.TaxMethod {
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
					idx--

					if idx < 0 || (currTransactionAmount > 0 && len(currencyDeposits) == 0 || (currTransactionAmount != 0 && len(currencyDeposits) == 1 && currencyDeposits[0].Amount == 0)) {
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
						CGL:         cgl - transaction.Gas,
					}
				}
			}
		}
	}
	return txTaxInfo, nil
}
