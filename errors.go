package txtax

import "errors"

var (
	ErrDepositTransactionWasNotFound = errors.New("Corresponding deposit transaction was not found\n")
	ErrNoCorrespondingDepositFound   = errors.New("No corresponding deposits found\n")
)
