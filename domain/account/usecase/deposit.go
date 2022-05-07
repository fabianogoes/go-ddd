package usecase

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-ddd/domain/account"
	"go-ddd/domain/transaction"
	"go-ddd/domain/transaction/model"
)

const (
	accountNotFound          = "account: %v not found to make deposit"
	updateAccountError       = "error to update account: %v on deposit"
	registerTransactionError = "registered a new transaction: %v"
)

type DepositUseCase struct {
	accountDomainService account.DomainService
	transactionService   transaction.DomainService
}

func NewDepositUseCase(ads account.DomainService, tds transaction.DomainService) *DepositUseCase {
	return &DepositUseCase{ads, tds}
}

func (u DepositUseCase) deposit(ctx context.Context, accountNumber int64, amount float64) error {
	accountFound, errFound := u.accountDomainService.FindAccountByNumber(ctx, accountNumber)
	if errFound != nil {
		log.Errorf(accountNotFound, accountNumber)
		return fmt.Errorf(accountNotFound, accountNumber)
	}

	if err := u.accountDomainService.UpdateAccount(ctx, accountFound.Deposit(amount)); err != nil {
		log.Errorf(updateAccountError, accountNumber)
		return fmt.Errorf(updateAccountError, accountNumber)
	}

	newTransaction := model.CreateNewTransaction(accountFound, amount, model.DEPOSIT)
	if t, err := u.transactionService.Register(newTransaction); err != nil {
		log.Errorf(registerTransactionError, accountNumber)
		return fmt.Errorf(registerTransactionError, accountNumber)
	} else {
		log.Infof("registered a new transaction: %v", t.ID)
		return nil
	}
}
