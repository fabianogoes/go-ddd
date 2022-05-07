package model

type TransactionType int64

const (
	DEPOSIT TransactionType = iota
	WITHDRAW
	PAYMENT
	TRANSFER
)

func (t TransactionType) Name() string {
	switch t {
	case DEPOSIT:
		return "DEPOSIT"
	case WITHDRAW:
		return "WITHDRAW"
	case PAYMENT:
		return "PAYMENT"
	case TRANSFER:
		return "TRANSFER"
	}
	return "unknown"
}
